Write-Output "|--------------------------------------------------------|"
Write-Output "| Creating Service Principal                             |"
Write-Output "|--------------------------------------------------------|"

$SUBSCRIPTION_ID = az account show `
    --query "id" `
    -o tsv

$SUBSCRIPTION_NAME = az account show `
    --query "name" `
    -o tsv

$SERVICE_PRINCIPAL_NAME = "github-action-$SUBSCRIPTION_ID"

$OUTPUT = az ad sp create-for-rbac `
    --name $SERVICE_PRINCIPAL_NAME `
    --sdk-auth `
    --skip-assignment

touch github_secret.json
Write-Output $OUTPUT > github_secret.json

Write-Output "|--------------------------------------------------------|"
Write-Output "| Assigning Roles to Service Principal                   |"
Write-Output "|--------------------------------------------------------|"

Start-Sleep -s 30

az role assignment create `
    --role "Owner" `
    --assignee http://$SERVICE_PRINCIPAL_NAME

az role assignment create `
    --role $ROLE `
    --assignee http://$SERVICE_PRINCIPAL_NAME

Write-Output "|--------------------------------------------------------|"
Write-Output "| Copy the contents of github_secret.json to your        |"
Write-Output "| GitHub organization or repository secret.              |"
Write-Output "|--------------------------------------------------------|"