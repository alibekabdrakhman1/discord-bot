# Discord Bot

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/your-discord-bot.git

2. Install dependencies
go get -u

3. Set up your Discord bot on the Discord Developer Portal and obtain your bot token.
4. Create a config.yam; file in the project config dir and add your bot token:
    ```
   Discord:
   Token: ""
5. go run cmd/main.go

## Commands

- **!help:**
  Help command.

- **!poll \<question\> | \<option1\> | \<option2\> | ...:**
  Create a poll.

- **!weather \<location\>:**
  Get current weather information.

- **!translate \<language\> \<text\>:**
  Translate text.

- **!remind \<time\> \<reminder\>:**
  Set a reminder.

- **!game rock|scissors|paper:**
  Play Rock-Paper-Scissors.
