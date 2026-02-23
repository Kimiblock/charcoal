# Netsock
A per-app firewall solution for Linux desktop. This is meant to be used by Portable. The name is derived from stockings and network. The socket is listened by systemd under `/run/netsock/control.sock`.

---

# Requests:

- `/add`: with a JSON which contains:

```
type IncomingSig struct {
	CgroupNested	string		// Nested control group path under /user.slice/user-1011.slice/user@1011.service
	RawDenyList	[]string	// JSON encoded list of raw rejected destinations. Accepts strings and IPs, special string private is interpreted as private IPs. Note that port 53 is always allowed.
	SandboxEng	string		// Sandbox engine ID
	AppID		string		// App ID
}
```

This returns a JSON containing:

```
type ResponseSignal struct {
	Success		bool
	Log		string
}
```