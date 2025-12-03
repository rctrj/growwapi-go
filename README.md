### Growwapi-go

Unofficial go wrapper for [Groww APIs](https://groww.in/trade-api/docs/curl). 
The documentation is directly picked from the API docs

Features:
- Strongly typed golang structs
- APIs exposed as golang methods

Future Features:
- Rate Limiting
- Retries

This started as a personal requirement, and I am adding modules when it's needed to me.
In case you want me to implement another module, let me know and it will be done

Below table shows the current status ( ✅: Done | ❔: May add | ❌: Probably won't add)

| Module                                        | Status |
|-----------------------------------------------|--------|
| Instruments                                   | ✅      |
| Orders                                        | ✅      |
| Smart Orders                                  | ❔      |
| Portfolio                                     | ❔      |
| Margin                                        | ❔      |
| Live Data                                     | ✅      |
| Historical Data (Deprecated, use Backtesting) | ❌      |
| Backtesting                                   | ✅      |
| Annexures                                     | ✅      |

### Installation
- Add this library to go mod
```
go mod github.com/rctrj/growwapi-go
```
  
### Contribution
Not accepting PRs at the moment since this is mostly a personal project at the moment, 
but open to the idea of opening the project for contributions if others are using this