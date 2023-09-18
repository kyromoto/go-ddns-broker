# Was soll mein System können?
- Ein (IP-)Endgerät [producer] (Router, Fritzbox, ...) bekommt eine neue IP vom ISP. Darauf hin soll dieses Endgerät dem System diese neue IP mitteilen. Das System soll dann diese IP an Drittsysteme [consumer] wie z.b. einen DNS Server oder eine Firewall weitergeben.
- Es soll pro Endgerät konfigurierbar sein welches Drittsystem über eine IP-Änderung informiert wird 
- Das System soll über eine Web-UI konfiguriert und überwacht werden können.
- Das System soll Mandantenfähig [client] sein. Das bedeutet, dass mehrere Endgeräte auf einen Mandant registriert sein können.