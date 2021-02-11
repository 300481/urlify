# urlify

Proxy-like Service which sends Push Notifications

## Environment

| Variable          | Description                                           |
|-------------------|-------------------------------------------------------|
| URL               | The URL of the data to respond                        |
| ALLOWED_REQUESTER | The IP of the system which is allowed to get the data |
| TOKEN             | The Pushover token                                    |
| USER              | The Pushover user                                     |

## Prerequisites

To use this service you need the following two prerequisites.

### Google Cloud Platform account

This code can be used with a GCP Cloud Function.

For that you need an account at [GCP](https://cloud.google.com/).

### Pushover Notifications account

This code needs an account at [Pushover Notifications](https://pushover.net/),

to send the notifications to your clients.

## Related Article

To get more information regarding the purpose of this code, please see my article on [Medium](https://dennis-riemenschneider.medium.com/how-to-encrypt-your-headless-linux-server-2de9c7f0f972)
