package exoscale

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnsexoscale "github.com/libdns/exoscale"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *libdnsexoscale.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.exoscale",
		New: func() caddy.Module { return &Provider{new(libdnsexoscale.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.APIKey = caddy.NewReplacer().ReplaceAll(p.Provider.APIKey, "")
	p.Provider.APISecret = caddy.NewReplacer().ReplaceAll(p.Provider.APISecret, "")

	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// exoscale {
//     api_key <string>
//     api_secret <string>
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_key":
				if d.NextArg() {
					p.Provider.APIKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "api_secret":
				if d.NextArg() {
					p.Provider.APISecret = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}

	if p.Provider.APIKey == "" {
		return d.Err("missing API Key")
	}

	if p.Provider.APISecret == "" {
		return d.Err("missing API Secret")
	}

	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)