from flask import Flask, render_template
import sqlite3

app = Flask(__name__)

@app.route('/')
def index():
    # Connectez-vous à la base de données SQLite
    conn = sqlite3.connect('ma_base_de_donnees.db')
    cursor = conn.cursor()

    # Exécutez une requête pour récupérer des données (exemple)
    cursor.execute('SELECT * FROM ma_table')
    data = cursor.fetchall()

    # Fermez la connexion à la base de données
    conn.close()

    # Rendez les données disponibles dans le modèle HTML
    return render_template('index.html', data=data)

if __name__ == '__main__':
    app.run(debug=True)
