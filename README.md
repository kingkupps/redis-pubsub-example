1. Install `websocat` with `brew install websocat`
1. Start the stack with `docker compose up -d`
1. Open a terminal window and run `websocat ws://localhost:1932/listen?receiver=D`
1. Open another terminal window and run `curl http://localhost:1932/send -X POST -d '{"receiver": "D", "message": "sup dawg"}'`

You should see the message from the last step appear in the websocket terminal window.
