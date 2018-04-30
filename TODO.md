### Todo for Dealer

Tracking 

 - [ ] search for all the available links in the domain and list all
 - [ ] persists keywords tagged with interval (e.g daily, instant, hourly)
- [ ] watches pages whether the keywords appear or not
   ```
	 for each keyword:
	     keyword.prio < page_link.last_call:
	         check whether the keyword exists in the page
	```
- [ ] calling every minutes for update
- [ ] take any link and check for keyword often
- [ ] take whole page hash it and persists it to avoid repeated keyword search
- [ ] while doing next check for the hash if change then do the keyword check 
- [ ] persists appeared keywords to queue for Notification 

Notification 
- [ ] check new events appeared from the queue server
- [ ] if new events appear to the msg queue(msg=push/email etc) https://github.com/rakanalh/scheduler
- [ ] email based on the update
		
