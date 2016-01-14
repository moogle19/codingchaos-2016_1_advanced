Aufgabe: Github Pullrequests
============================

Ziel dieser Aufgabe ist es alle offenen Pullrequests eines Github-Repositories
in folgender Form aufzulisten:

    <request-id>  <username-author> <commit-message>

Beispiel:

    ./pullrequests.py mkdocs mkdocs                                                                                                                            (git)-[master] [0]

    781           ksimka  Fix readthedocs theme subnav padding
    772        jimporter  Change the styling of inline code and allow links to colorize inline …
    752           dmehra  Make 'Edit on GitHub' link point at individual pages. Addresses issue #269.
    ...


Programmaufruf
--------------

Der Aufruf der Anwendung soll folgendermaßen aussehen:

    ./pullrequests <github-user> <github-repo>


Tipps
-----

- Sinnvollerweise sollte die Gihub-API genutzt werden
- Githubs RESTful-API ist zu erreichen unter:
        https://api.github.com/
- Die API liefert alle Daten als JSON zurück
