**DEVELOPER INSTRUCTIONS:**

- Update module name in go.mod
- Update dependencies to latest versions
- Update name and year in license
- Customize configuration and Caddyfile parsing
- Update godocs / comments (especially provider name and nuances)
- Update README and remove this section

---

Exoscale module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [Exoscale](https://www.exoscale.com/).

## Caddy module name

```
dns.providers.exoscale
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "exoscale",
				"api_key": "{env.EXOSCALE_API_KEY}",
                "api_secret": "{env.EXOSCALE_API_SECRET}"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns exoscale {
	    api_key {$EXOSCALE_API_KEY}
	    api_secret {$EXOSCALE_API_SECRET}
	}
}
```

```
# one site
tls {
	dns exoscale {
	    api_key {$EXOSCALE_API_KEY}
	    api_secret {$EXOSCALE_API_SECRET}
	}
}
```