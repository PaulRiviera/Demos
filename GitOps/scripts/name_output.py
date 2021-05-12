import sys

version = sys.argv[1]
azureRegion = sys.argv[2]
environment = "demo"
owner = "msft"

possibleAzureRegions = {
    "eastus":"eu",
    "eastus2":"eu2",
    "westus":"wu",
    "westus2":"wu2"
    }

def GetAzureShortRegion(regionName):
    for key in possibleAzureRegions.keys():
        if key == regionName:
            return possibleAzureRegions[key]

def OutputResourceGroupName(owner, version, environment):
    print("::set-output name=resource_group_name::rg-{}-{}-{}".format(owner, version, environment))

def OutputKeyVaultName(owner, version, environment):
    print("::set-output name=key_vault_name::kv-{}-{}-{}".format(owner, version, environment))

def OutputAzureContainerRegistry(owner, version, environment):
    print("::set-output name=registry_name::acr{}{}{}".format(owner, version, environment))

def OutputResourceGroupName(owner, version, region, environment):
    print("::set-output name=resource_group_name::rg-{}-{}-{}-{}".format(owner, version, region, environment))

def OutputClusterName(owner, version, region, environment):
    print("::set-output name=cluster_name::aks-{}-{}-{}-{}".format(owner, version, region, environment))

def OutputLogWorkspaceName(owner, version, region, environment):
    print("::set-output name=log_workspace_name::logs-{}-{}-{}-{}".format(owner, version, region, environment))

shortRegionName = GetAzureShortRegion(azureRegion)

OutputResourceGroupName(owner, version, environment)
OutputKeyVaultName(owner, version, environment)
OutputAzureContainerRegistry(owner, version, environment)
OutputResourceGroupName(owner, version, shortRegionName, environment)
OutputClusterName(owner, version, shortRegionName, environment)
OutputLogWorkspaceName(owner, version, shortRegionName, environment)
