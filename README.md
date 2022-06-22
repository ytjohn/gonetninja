# Welcome to Net Ninja

Net Ninja lets amateur radio net control operators record net times, checkins, and comments.

## Development notes

- /nets - page displaying nets
- /activity - page displaying net activity, most commonly filtered by net id


## tables

Nets

- id (uuid)
- name
- planned start
- planned finish

```sqlite
sqlite> insert into nets values (uuid(), strftime('%Y-%m-%d %H-%M-%S','now'), strftime('%Y-%m-%d %H-%M-%S','now'), "test two", strftime('%Y-%m-%d %H-%M-%S','now'), strftime('%Y-%m-%d %H-%M-%S','now'));
sqlite> insert into nets values (uuid(), strftime('%Y-%m-%d %H-%M-%S','now'), strftime('%Y-%m-%d %H-%M-%S','now'), "test three", strftime('%Y-%m-%d %H-%M-%S','now'), strftime('%Y-%m-%d %H-%M-%S','now'));
sqlite> select * from nets;
c31911fc-5d3c-4b18-b4b1-1e081aa6effd|2022-06-20 23-47-29|2022-06-20 23-47-29|test one|2022-06-20 23-47-29|2022-06-20 23-47-29
ccad3aad-c9ea-4891-a604-8d02e0968ce8|2022-06-20 23-47-47|2022-06-20 23-47-47|test two|2022-06-20 23-47-47|2022-06-20 23-47-47
66b685a9-ea20-4a14-b766-4d23f362be4b|2022-06-20 23-48-17|2022-06-20 23-48-17|test three|2022-06-20 23-48-17|2022-06-20 23-48-17
```


Net Activity

- id (uuid)
- created (auto)
- updated (auto)
- netid (foreign key)
- entered_by (username)
- time_at (manual, with suggestion by frontend ui)
- action
  - open
  - assign (assign net control)
  - checkin
  - checkout
  - comment
  - close
- callsign
- details



By making it an activity trail, we can have people check in before the net opens
(early checkins), open the net, close the net, check poeple out of the net early,
add comments, change net control

## Database Setup

It looks like you chose to set up your application using a database! Fantastic!

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

You will also need to make sure that **you** start/install the database of your choice. Buffalo **won't** install and start it for you.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started your database, now Buffalo can create the databases in that file for you:

```console
buffalo pop create -a
```

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

```console
buffalo dev
```

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

Good luck!

[Powered by Buffalo](http://gobuffalo.io)
