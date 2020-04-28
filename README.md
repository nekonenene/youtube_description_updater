# YouTube Description Updater

You can update descriptions of all your YouTube videos via CLI.


## Required

You must put `client_secret.json` in the working directory.

Log in to your [Google Developers Console](https://console.developers.google.com) account and create the OAuth client ID.  
See here: [https://support.integromat.com/hc/en-us/articles/360025257393-Connecting-YouTube-to-Integromat-via-Google-OAuth-Client](https://support.integromat.com/hc/en-us/articles/360025257393-Connecting-YouTube-to-Integromat-via-Google-OAuth-Client)

After creating the OAuth client ID, download the JSON file and rename to `client_secret.json`.

<img src="doc/download_json.png">


## Usage

When you want to update an URL in descriptions of all your videos:

```sh
youtube_description_updater -target-string https://old.url.com -replacement-string https://new.url.com
```

You can see all options:

```sh
youtube_description_updater -h
``` 


## Caution

Usually, [YouTube Data API v3 quota](https://developers.google.com/youtube/v3/getting-started#quota) limit per day is 10,000, and a write operation spends around 50 quota. As a result, if your channel manages over 200 videos, this command can not finish successfully. Please check `-limit` option and set the everyday cron job.
