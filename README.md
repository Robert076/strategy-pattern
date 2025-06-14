# 🌎 strategy-pattern
Trying to get used to the strategy design pattern, hence this practice repo. Followed tutorial on refactoring.guru.

This design pattern is recommended to be used in the case of long if statements, where there are ifs after ifs. For example in my message code if I had the Message struct as:

```
type Message struct {
	text         string
	sendstrategy string
}
```

instead of

```
type Message struct {
	text         string
	sendstrategy sendstrategy
}
```


Then every time I sent a message I had to do:

```
*pseudocode*
if message.sendstrategy == "sms"
  smsservice.send(message)
else if message.sendstrategy == "carrierpigeon"
  carrierpigeonservice.send(message)

etc.
```

You notice how scaling the program gets increasingly harder as each new method adds another if statement.

But if instead we use a strategy no new code is needed except for the implementation of the new strategy, and it can be then passed down if it implements the strategy interface.

Making the final code just:

```
Message.send()
```

That's it!

Very useful pattern.
