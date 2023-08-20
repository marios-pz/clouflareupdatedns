# clouflareupdatedns
Update DNS Records easily

# Instructions

Before running the script you need to setup 
environmental variables that are required to authenticate to cloudflare
while also specifying which dns record to change.

```
export ZONE_ID=""         # Cloudflare Zone ID for the DNS records
export AUTH_TOKEN=""      # Authentication token for accessing the Cloudflare API
export DNS_NAME=""        # The name of the DNS record you want to update
export DNS_RECORD_ID=""   # The specific DNS record ID you want to update
export PUBLIC_IP=""       # The new public IP address to set for the DNS record
```

