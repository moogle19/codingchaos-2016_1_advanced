Aufgabe: Timing Attack
======================

Ziel dieser Aufgabe ist es das Passwort des Programmes `./crypto`
herauszufinden indem eine Timing-Attacke [1] durchgeführt wird.


Einleitung
----------

Aufgerufen wird das Programm `./crypto` mit

     ./crypto <password>

Es hat die folgende Funktion:

 - Ist das Passwort richtig, hat das Programm `./crypto` einen Returncode 0.
 - Ist das Passwort falsch, hat das Programm `./crypto` einen Returncode 1.

Die Programmierer von `./crypto` haben bei der Implementierung mehrere Fehler
gemacht:

 - Wird ein Passwort übergeben, das länger ist als das eigentliche Passwort,
   so geht das Programm mit dem Returncode `2` kaputt.
 - Für jeden richtigen Buchstaben im Passwort braucht `./crypto` länger bis es
   sich beendet.


Ziel
----

Schreibe ein Programm, das eine Timing-Attacke auf das Programm `./crypto`
ausführt und dadurch das Passwort herausfindet. Hierfür genügt es das Programm
`./crypto` per System-Call [2] mit verschiedenen Parametern aufzurufen und die
benötigte Zeit zu stoppen. Dadurch kann das Passwort herausgefunden werden.

Es darf davon ausgegangen werden, dass alle Passwörter nur aus Großbuchstaben
bestehen.

Programmausgabe
---------------

Dein Programm sollte am Ende der Berechnung das richtige Passwort in der
Konsole ausgeben.


[1] https://en.wikipedia.org/wiki/Timing_attack

[2] http://linux.die.net/man/3/system https://docs.python.org/2/library/os.html#os.system http://docs.oracle.com/javase/7/docs/api/java/lang/Runtime.html#exec%28java.lang.String%29
