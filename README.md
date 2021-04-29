# channel-alarm
 Listen on a channel to trigger a function

 This is just a little piece of test code for my larger alarm clock build. This program has multiple
 goroutines sending to a channel, and a main function listening to the channel. Depending on what 
 goroutine sends, the main function executes different responses. This is to test having multiple alarms
 going off at different times, sometimes even interrupting each other (which we'll eventually need the
 context package for)

 the main issue with the code as it stands is that if all alarms are turned off, there will be a problem
 with deadlock, unless the channel is closed, in which case main will exit. The solution to this is most likely to have all of these things in tools, and have the main web server running and calling
 these functions from tools as needed. There will need to be some sort of check() function to see if
 any alarms are running, since if there are not then listen() will need to be called again
