# salat-time-bot

This bot is intended to run on my personal machine and ping a telegram bot 5 minutes before and after a prayer time has begun

[v1] features

- Fetch prayer times from a free API (undetermined) on system login
  - If no internet; loopback [v2]
- Check every `n` minutes to see if current time within 5 minutes threshold of any prayer and send telegram ping with prayer name
  - Make PUB/SUB or something similar so that just before and after five minutes a ping is sent to avoid continuous checks [v2]
