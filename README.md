## purpose

periodically sends a get request to a given url.  Google onhub has a bug where it forgets about a device that it is forwarding too if that device is not actively reaching out to the internet.  so deploying this daemon ensures the given device will always be pinging the internet and prevents the forwarding port from blocking.