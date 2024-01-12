$(function () {
    const socket = io();

    // Écouter les messages du serveur
    socket.on('chat message', (msg) => {
        $('#messages').append($('<li>').text(msg));
    });

    // Envoyer un message lorsque le formulaire est soumis
    $('form').submit(function() {
        socket.emit('chat message', $('#m').val());
        $('#m').val('');
        return false;
    });
});
