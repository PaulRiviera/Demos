param(
    [Parameter(Mandatory=$true)]
    [String]
    $RESOURCE_GROUP_NAME,
    [Parameter(Mandatory=$true)]
    [String]
    $REGION_NAME,
    [Parameter(Mandatory=$true)]
    [String]
    $KEY_VAULT_NAME,
    [Parameter(Mandatory=$true)]
    [String]
    $AZURE_CONTAINER_REGISTRY_NAME,
    [String]
    $AZURE_CONTAINER_REGISTRY_SKU = "Basic",
    [Parameter(Mandatory=$true)]
    [String]
    $DEMO_NAME
)

az group create `
    --name $RESOURCE_GROUP_NAME `
    --location $REGION_NAME

az keyvault create `
    --resource-group $RESOURCE_GROUP_NAME `
    --name $KEY_VAULT_NAME `
    --location $REGION_NAME `
    --sku standard

az acr create `
    --resource-group $RESOURCE_GROUP_NAME `
    --location $REGION_NAME `
    --name $AZURE_CONTAINER_REGISTRY_NAME `
    --sku $AZURE_CONTAINER_REGISTRY_SKU `
    --admin-enabled
    