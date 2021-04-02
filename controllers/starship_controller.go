/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	starshipsv1beta1 "github.com/julienstroheker/kubebuilder2-starwars-example/api/v1beta1"
)

// StarshipReconciler reconciles a Starship object
type StarshipReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=starships.starwars.julienstroheker.com,resources=starships,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=starships.starwars.julienstroheker.com,resources=starships/status,verbs=get;update;patch

func (r *StarshipReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("starship", req.NamespacedName)

	instance := &starshipsv1beta1.Starship{}
	err := r.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return ctrl.Result{}, err
	}

	var foundShip shipResponse
	var found bool

	url := "https://swapi.dev/api/starships/"

	for !found {
		data := getData(url)

		ship := findShip(data.Results, instance.Spec.Name)

		if ship != nil {
			foundShip = *ship
			found = true
		}

		if ship == nil && data.Next != "" {
			url = data.Next
		}

		if !found && data.Next == "" {
			return ctrl.Result{}, errors.NewBadRequest("Invalid Ship Name given")
		}
	}

	instance.Spec.Name = foundShip.Name
	instance.Status.Capacity = foundShip.Capacity
	instance.Status.Costs = foundShip.Costs
	instance.Status.Crew = foundShip.Crew
	instance.Status.Model = foundShip.Model
	instance.Status.Passengers = foundShip.Passengers

	err = r.Update(context.TODO(), instance)

	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

type apiResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous,"`
	Results  []shipResponse `json:"results"`
}

type shipResponse struct {
	Name         string `json:"name"`
	Model        string `json:"count"`
	Manufacturer string `json:"manufacturer"`
	Costs        string `json:"cost_in_credits"`
	Passengers   string `json:"passengers"`
	Crew         string `json:"crew"`
	Capacity     string `json:"cargo_capacity"`
}

func getData(url string) apiResponse {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	var result apiResponse
	err = json.Unmarshal(data, &result)

	return result
}

func findShip(ships []shipResponse, shipname string) *shipResponse {
	for _, ship := range ships {
		if ship.Name == shipname {
			return &ship
		}
	}
	return nil
}

func (r *StarshipReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&starshipsv1beta1.Starship{}).
		Complete(r)
}
