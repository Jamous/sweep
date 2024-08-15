Sweep
=====
Sweep will ping every IPv4 address in every active network directly connected to a device, to include VPN tunnels. Sweep is useful for network discovery and filling out IPv4 arp tables. It does this by sending one ICMP echo request, then closing the connection. This program will set off some antivirus programs.

Compiled binaries are located under bin/

Sweep implements Pingsweep, which you can see here. https://github.com/Jamous/pingsweep

Video
https://www.youtube.com/watch?v=SvIUSOy-Suc

Options
-------
You can specify the network you want to ping in CIRD notation.

```
./sweep 192.168.10.0/24
```

Versions
--------
v0.0.1 - initial release - 08/09/24
