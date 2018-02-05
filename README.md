# Simple DB

The description of the assignment can be found [here](https://render.githubusercontent.com/view/ipynb?commit=51186bebc8cda97b5d150e87eb11633e39f75611&enc_url=68747470733a2f2f7261772e67697468756275736572636f6e74656e742e636f6d2f646174736f66746c796e6762792f736f667432303138737072696e672d6461746162617365732d7465616368696e672d6d6174657269616c2f353131383662656263386364613937623564313530653837656231313633336533396637353631312f6c6563747572655f6e6f7465732f30312d496e74726f5f746f5f44422e6970796e62&nwo=datsoftlyngby%2Fsoft2018spring-databases-teaching-material&path=lecture_notes%2F01-Intro_to_DB.ipynb&repository_id=118117359&repository_type=Repository#Assignment-1---Simple-DB-with-Hashmap-based-Index).

I decided to get a little bit out of my comfort zone, and write this little project in [Go](https://golang.org/).

## How it works

The idea was to build a very simple database with a hashmap-based index. I've decided to use the Go programming language, along with the [Bolt](https://github.com/boltdb/bolt) library for this task.

The idea is this: whenever you launch the program and give it a command along with some values, a series of events will happen:

1. The binary database file is read. A new file will be created if it doesn't exist.
2. The command and value(s) are read.
3. The input command is executed by operating on the database file with Bolt's commands.
4. If the set-commad is executed, the input values will be encoded into bytes. If the get-command is executed, returned values are decoded into strings and returned. 

## How to use it

### Prerequisites

You need to have Go installed on your machine in order to run this project. That can be done from [here](https://golang.org/doc/install).

Alternatively, you can launch a [Docker](https://www.docker.com/) container or a [Vagrant](https://www.vagrantup.com/) virtual machine with Go pre-installed.

If you're using Windows, it's recommended to run the program with the [Git Bash](https://git-scm.com/downloads) terminal, or any other Linux-based terminal of your choice.

### Mac/Linux/Windows

1. Clone this repository in a location of your choice:

   ```````bash
   $ git clone https://github.com/dbdness/simple_db.git
   ```````


2. Build the project:

   ```bash
   $ go build
   ```

3. Place a value along with a key in the database using the `--set` command and its values. The first argument after the command is the key, the other is the value:

   ```bash
   $ ./simple_db --set 1 Test
   ```

4. Retrieve the value of the key you've just placed in the database, by using the `—-get` command:

   ``` bash
   $ ./simple_db --get 1
   ```

5. To delete a key-value pair, use the `—-delete` command along with the key of the entry you want to delete:

   ```bash
   $ ./simple_db --delete 1
   ```