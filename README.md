# charcoal
A filter for sandboxes

---

Environment variables:

- `RUNTIME_DIRECTORY`: specify where charcoal should place the control socket in. A sensible default is `/run/charcoal`, set by systemd. Which makes charcoal listen on `/run/charcoal/control.sock`.
	* Downstream apps may not support different socket path