/**
 * @author Yamao Cuzou <yamao.cuzou@ynov.com>
 * @file Description
 * @desc Created on 2024-01-10 11:14:02 pm
 * @copyright Cuzou Corporation
 */

function init_client(fonction, a, b, nombreRectangles) {
    const largeurRectangle = (b - a) / nombreRectangles;
    let somme = 0;

    for (let i = 0; i < nombreRectangles; i++) {
        const xGauche = a + i * largeurRectangle;
        const xDroite = xGauche + largeurRectangle;
        const hauteurRectangle = (fonction(xGauche) + fonction(xDroite)) / 2;
        somme += hauteurRectangle * largeurRectangle;
    }
    return somme;
}

function init_server(n, k) {
    function factorielle(nombre) {
        if (nombre === 0 || nombre === 1)
        {
            return 1;
        } else
        {
            return nombre * factorielle(nombre - 1);
        }
    }
    if (k < 0 || k > n)
    {
        return 0;
    }
    const coefficient = factorielle(n) / (factorielle(k) * factorielle(n - k));
    return coefficient;
}