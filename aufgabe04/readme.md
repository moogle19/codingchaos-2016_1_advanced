Aufgabe: Waveform für Audio-Datei
=================================

Ziel dieser Aufgabe ist die grafische Darstellung einer vorgegebenen
Audiodatei.

Beispiel:

![Waveform example](waveform.png)

 - Eingabe ist eine vorgegebene Binärdatei mit Audio-Rohdaten
 - Ausgabe ist eine Bilddatei im [PBM-Format
   ](https://en.wikipedia.org/wiki/Netpbm_format)
 - Bildgröße soll 1000x255 Pixel sein

Der visuelle Effekt ist wichtig. Es kommt nicht allzu sehr auf Genauigkeit an.


Programmaufruf
--------------

Der Aufruf der Anwendung soll folgendermaßen aussehen:

    ./waveform input.data output.pbm


Vorgaben
--------

Gegeben ist eine Audiodatei mit:

 - Einem Kanal
 - Die Datei enthält ein Integer/Sample-Array
 - Sample-Format: Signed, 16bit, lower endian (s16le)


Tipps
-----

*(bzw. was muss ich tun)*

 - Binärdaten einlesen
 - Daten vertikal skalieren: Werte müssen zusammengefasst werden um auf eine
	Gesamtbreite von 1000 Pixeln zu kommen. Die Art der Zusammenfassung ist
        in der Bewertung nicht entscheident.
 - Daten horizontal skalieren: Der Ausschlag nach oben muss berechnet werden.
	Der Maximalwert eines Samples ist `2^16`. Das Bild ist aber nur 255 Pixel
	hoch.
 - Bilddatei schreiben. Es darf binäres und nicht-binäres PBM-Format genutzt werden.


Eigene Audiodaten
-----------------

Eigene Audiodaten können leicht mit Hilfe von [FFmpeg](http://ffmpeg.org) in
das vorgegebene Format gebracht werden:

    ffmpeg -i in.wav -map 0:a -c:a pcm_s16le -ac 1 -f data out.data

Das Eingabeformat ist FFmpeg dabei recht egal.
