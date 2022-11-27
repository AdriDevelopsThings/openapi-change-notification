# openapi-change-notification

You are using a lot of APIs with OpenAPI specification and want to get notifications by email when an endpoint which you are using gets deprecated? openapi-change-notification can send you these notifications.
All subscribed API endpoints are fetched every hour. If there is a change of an endpoint you subscribed, it sends you a notification by email. 

# How to use?

Go to [https://openapi.adridoesthings.com](https://openapi.adridoesthings.com) or host your own instance with docker or run it yourself (you can build a binary by typing in ``go build .``.

## Environment
You need the following environment variables:
```
REDIS_URL=localhost:6379 # yes you need a redis server cause of caching
FRONTEND_URL=https://the-url-of-the.frontend.com # the public url of your instance
SMTP_SERVER=mail.example.org
SMTP_PORT=587
SMTP_USERNAME=username
SMTP_PASSWORD=password
SMTP_FROM_ADDRESS=noreply@example.org
HCAPTCHA_SECRET=YOUR_HCAPTCHA_SECRET # hcaptcha is in use because we want to avoid people from spamming with the subscription form
```

Also you need to save the `/db.sql` as a volume in docker or change the `SQLITE_PATH` environment variable to another path.

# Subscribe
You can subscribe to new OpenAPI configurations, paths and methods by filling the form on the index site of the frontend. Every subscription will be merged, so you are not able to create duplicates. Also, you will receive just one email verification email. The OpenAPI configuration is checked while you are subscribing, so that an error will be shown if the fetch was not successful.
After you have subscribed, you get an email verification link by email. You must open this link, or you won't get notifications.

# Unsubscibe
If you want to unsubscribe from all notifications on openapi-change-notification go to the frontend URL (public instance: (https://openapi.adridoesthings.com/unsubscribe)[https://openapi.adridoesthings.com/unsubscribe]) with `/unsubscribe` and enter your email address. An unsubscription verification email will be sent to you, so you have to click on the link to finally unsubscribe from all notifications.
