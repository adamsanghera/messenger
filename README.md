# Messenger

## Workers

1. DB Update Worker

2. Conversation List Update Worker --> Most complicated

## Data Structs

1. User
  * Name
  * UID
  * CIDs [Stored as linked list]

2. Conversation
  * Messages [Stored as linked list]
    * Author
    * Timestamp
    * Content
  * UserIDS [Stored as slice]
  * CID