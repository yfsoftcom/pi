from myio import load_gpio, on, off, is_setup
from flask import Flask, render_template, jsonify
app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/toggle/<channel>/<value>')
def toggle(channel, value):
    
    if not load_gpio:
        return jsonify({'code': 0})
    if not is_setup(channel):
        return jsonify({ 'code': -1 })
    if int(value) == 1:
        on(channel)
    else:
        off(channel)

    return jsonify({'code': 0}) 

if __name__ == '__main__':
    app.run(debug=True, port=8000, host='0.0.0.0')