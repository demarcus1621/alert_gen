from flask import Flask, request

app = Flask(__name__)

@app.route('/api/v2/passwords')
def send_password_list():
    pw_list = [
        "matthew",
        "robert",
        "danielle",
        "forever",
        "family",
        "jonathan",
        "987654321",
        "computer",
        "whatever",
        "dragon",
        "vanessa",
        "cookie",
        "naruto",
        "summer",
        "sweety",
        "spongebob",
        "joseph",
        "junior",
        "softball",
        "taylor",
        "yellow",
        "daniela",
        "lauren",
        "mickey",
        "princesa",
        "alexandra",
        "alexis",
        "jesus",
        "estrella",
        "miguel",
        "william",
        "thomas",
        "beautiful",
        "mylove",
        "angela",
        "poohbear",
        "patrick",
        "iloveme",
        "sakura",
        "adrian",
        "alexander",
        "destiny",
        "christian",
        "121212",
        "sayang",
        "america",
        "dancer",
        "monica",
        "richard",
        "112233"
    ]
    return pw_list

@app.route('/api/v2/receive_data')
def receive_data():
    pass
    # data = json.loads(request.body)

@app.route('/api/v2/receive_file')
def receive_file():
    pass

@app.route('api/v2/send_command')
def send_command():
    pass

@app.route('api/v1/send_command')
def send_random():
    pass
