import flask

application = flask.Flask(__name__)
application.config['SECRET_KEY'] = 'SecretKeyHere'


@application.errorhandler(Exception)
def errorpage(Exception):
    return flask.render_template("errorpage.html",error = Exception)

@application.route("/")
def index():
    return flask.render_template("index.html")

@application.route("/about")
def about():
    return flask.render_template("about.html")

if __name__ == '__main__':
    application.run(debug = False)