# Message

AA | BB BB BB BB | CC 

-> AA: Type of message
-> BB: Size of the message
-> CC: Message 

# Message Type

Client -> Server
00 - Ping
01 - Auth
02 - Join Channel
03 - Leave Channel

Client -> Server
10 - Client Update Health 
11 - Client Update Inventory
12 - Client Update ItemBox
13 - Client Update Selected Inventory
14 - Client Update Last Item Seen
15 - Client Update Room
16 - Client Update Enemy

Server -> Client
20 - Server Update Health
21 - Server Update Inventory -> (RE2 Only)
22 - Server Update ItemBox 

# Message Structure

XX XX XX XX | YY

-> XX: Message ID
-> YY: Message Content