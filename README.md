# channel-alarm
 Listen on a channel to trigger a function

 This is just a little piece of test code for my larger alarm clock build. This program has multiple
 goroutines sending to a channel, and a main function listening to the channel. Depending on what 
 goroutine sends, the main function executes different responses. This is to test having multiple alarms
 going off at different times, sometimes even interrupting each other (which we'll eventually need the
 context package for)

 the main issue with the code as it stands is that if all alarms are turned off, there will be a problem
 with deadlock, unless the channel is closed, in which case main will exit

 NOTE: There is currently a bug where sometimes the alarm message will send *after* the alarm has been
 turned off. Probably not a huge deal, and seems caused by the fact that the message is being sent on
 a channel, whereas the statement that the alarm has been turned off is not
