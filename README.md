# Signal Vip Bot

A Telegram bot that get an Excel file and sent it to A Channel.

## Requirements:
- MongoDB

### Change access control settings
- there can be DB_ACCESS_CONTROL in .env file for mongodb access control settings, default state is on.
### Project tree
```bash
Signal Vip Bot
└───src
    ├───cmd
    │   ├───cli
    │   └───providers
    ├───internal
    │   ├───data
    │   │   ├───datasource
    │   │   │   ├───mongo
    │   │   │   │   └───models
    │   │   │   └───telegram
    │   │   └───repositories
    │   │       └───mocks
    │   ├───domain
    │   │   ├───authing
    │   │   ├───entities
    │   │   ├───mocks
    │   │   └───publishing
    │   └───presentation
    │       ├───cli
    │       └───telegram
    │           ├───controllers
    │           ├───helpers
    │           └───structs
    ├───pkg
    │   ├───daterefactor
    │   ├───excel
    │   └───validation
    └───statics
```