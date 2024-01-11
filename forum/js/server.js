const express = require('express');
const http = require('http');
const socketIo = require('socket.io');

const app = express();
const server = http.createServer(app);
const io = socketIo(server);

app.use(express.static('public'));

io.on('connection', (socket) => {
    console.log('Un utilisateur s\'est connecté');

    // Écouter les messages du client
    socket.on('chat message', (msg) => {
        // Diffuser le message à tous les clients connectés
        io.emit('chat message', msg);
    });

    // Gérer la déconnexion de l'utilisateur
    socket.on('disconnect', () => {
        console.log('Un utilisateur s\'est déconnecté');
    });
});

const PORT = process.env.PORT || 3000;

server.listen(PORT, () => {
    console.log(`Serveur en écoute sur le port ${PORT}`);
});
