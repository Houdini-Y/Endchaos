# EndChaos

**The last recon tool you'll ever need.**

One command. Every endpoint. Full verb matrix. Crowdsourced intelligence.

```bash
endchaos uber.com --full --contribute
# → 187,432 endpoints with methods, tech, titles, interesting flags
# → automatically contributes to central database
```
## EndChaos combines:

	- All passive sources (Wayback, CommonCrawl, URLScan, OTX, Chaos, crt.sh)
	- Katana crawling + XHR/JS extraction
	- httpx with full verb testing
	- Kiterunner + aggressive bruteforce
	- Auto-detection of GraphQL, WebSocket, uploads, debug panels
	- Local cache + diff mode
	- Crowdsourced endpoint intelligence (opt-in)

This is the tool that ends the "8-hour recon" era.

Built on: ProjectDiscovery tools + blood, sweat, and hatred for slow recon

## Installation
```bash
go install -v github.com/endchaos-project/endchaos/cmd/endchaos@latest
```
or download binary from Releases

## Quick start
```bash
endchaos hackerone.com --full --contribute
```

Join the movement. Star. Contribute. Break the internet.
