# Creating / Deleting CR every 10 sec to generate API Activity
 
while(1)
{
   Write-Host "Apply CR"
   kubectl apply -f ./config/samples/starships_v1beta1_starship.yaml
   Write-Host "Sleep 10 Sec"
   start-sleep -seconds 10
   Write-Host "Delete CR"
   kubectl delete -f ./config/samples/starships_v1beta1_starship.yaml
   Write-Host "Sleep 10 Sec"
   start-sleep -seconds 10
} 
