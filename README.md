# Slack File Upload Bot

A simple Go application that uploads files to a specified Slack channel using the Slack API.

## Prerequisites

- Go 1.16 or higher
- A Slack workspace where you have permission to create apps
- A Slack bot token with appropriate permissions

## Setup Instructions

1. **Create a Slack App**
   - Go to [api.slack.com/apps](https://api.slack.com/apps)
   - Click "Create New App"
   - Choose "From scratch"
   - Name your app and select your workspace

2. **Configure Bot Permissions**
   - In your app settings, go to "OAuth & Permissions"
   - Under "Scopes", add these bot token scopes:
     - `files:write`
     - `chat:write`
   - Install the app to your workspace
   - Copy the "Bot User OAuth Token" (starts with `xoxb-`)

3. **Get Channel ID**
   - In Slack, right-click on the channel you want to upload files to
   - Select "Copy link"
   - The last part of the URL is your channel ID (starts with `C`)

4. **Set Up Environment Variables**
   - Create a `.env` file in the project root
   - Add your credentials:
     ```
     SLACK_BOT_TOKEN=xoxb-your-token-here
     CHANNEL_ID=your-channel-id-here
     ```

5. **Install Dependencies**
   ```bash
   go mod download
   ```

6. **Run the Application**
   ```bash
   go run main.go
   ```

## Project Structure

- `main.go` - Main application code
- `.env` - Environment variables (not tracked in git)
- `go.mod` - Go module file
- `go.sum` - Go module checksums

## How It Works

The application:
1. Loads environment variables from `.env`
2. Connects to Slack using your bot token
3. Uploads the specified file to the designated channel
4. Prints the upload status

## Troubleshooting

- If you get a "token invalid" error, check your `SLACK_BOT_TOKEN`
- If you get a "channel not found" error, verify your `CHANNEL_ID`
- Make sure your bot has been invited to the channel
- Ensure the file you're trying to upload exists in the project directory

## Security Notes

- Never commit your `.env` file
- Keep your bot token secure
- The `.gitignore` file is set up to prevent accidental commits of sensitive data

## License

MIT License 