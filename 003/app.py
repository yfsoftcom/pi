from myio import load_gpio, on, off
from flask import Flask, render_template, jsonify
app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/toggle/<ionumber>/<value>')
def toggle(ionumber, value):
    print(ionumber, value)
    if not load_gpio:
        return jsonify({'code': 0})
        
    if int(value) == 1:
        on(ionumber)
    else:
        off(ionumber)

    return jsonify({'code': 0}) 

if __name__ == '__main__':
    app.run(debug=True, port=8000, host='0.0.0.0')