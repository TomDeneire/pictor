# Discovering IIIF Manifests with Pictor

![screenshot.png](screenshot.png)

## Idea

Discovering [IIIF](https://iiif.io/) resources can be hard.

Although the protocol does specify a dedicated [Discovery API](https://iiif.io/api/discovery/1.0/) it is not often implemented by institutions. (At [Anet](https://anet.be) we are guilty of the same). Moreover, this API has no straightforward way to obtain a full collection. It is certainly not as straightforward as with [OAI-PMH](https://www.openarchives.org/pmh/) for instance, that offers the verb `ListIdentifiers`.

The IIIF documentation does have an interesting [Guide to finding IIIF resources](https://iiif.io/guides/finding_resources/) which features a list of IIIF collections.

With that information I was able to scrape several of these collections and aggregate them into a corpus of close to 2,000,000 IIIF manifests. The resulting lists are available in this repository.

Currently, it features manifests of the following institutions / collections:

- [Anet library network](https://www.uantwerpen.be/en/projects/anet/)
- [Bayerische Staatsbibliothek (BSB) / Munich Digitization Centre (MDZ)](https://www.digitale-sammlungen.de/en/)
- [University of Toronto](https://collections.library.utoronto.ca/)
- [Digital Commonwealth](https://digitalcommonwealth.org/)
- [Getty Institute](https://iiif.io/guides/guides/search.getty.edu/)
- [Wikidata](https://www.wikidata.org/)

This repository has two purposes. One it offers a place to **store IIIF collections** and make them available for others. Two it uses those collections to host a **discovery tool** with a sample of them:

## Harvesting

Harvesting the IIIF manifests was done with Python scripts that harvest these links in a variety of ways.

Some institutions, like the [Bayerische Staatsbibliothek](https://www.digitale-sammlungen.de/en/) or the [University of Toronto](https://collections.library.utoronto.ca/) let you scrape collections from their Presentation API (although not without a certain limit, unfortunately). Others, like [Digital Commonwealth](https://digitalcommonwealth.org/) have OAI-PMH that get you the necessary identifiers. Still others, like the [Getty Institute](https://iiif.io/guides/guides/search.getty.edu/) or [Wikidata](https://www.wikidata.org/) offer a SPARL endpoint.

I harvested all manifests I could find for the repository and also made random 5,000 manifest samples of the collections for the discovery tool. For this, good old Unix tools are still amazingly good:

```bash
shuf digitalcommonwealth.txt | head -n 5000 > digitalcommonwealth_sample.txt
```

## Indexing

Requesting and indexing the IIIF manifests was done with a Go script (since Go is really strong for concurrency) and the result was piped as triple statements from stdout to a plain textfile. (I found this a handy alternative to having to set up a database like PostgreSQL or something similar that could handle concurrent writing).

The resulting triple store was then turned into a number of JSON files, including one for the IIIF manifest identifiers and their matching sequential number. I used base-85 numbers for the latter, as this gave me a very efficient way to encode large numbers.

In total, this process only takes a couple of hours for the current sample of ca. 30,000 manifests.

## Web application

Finally, a web interface with some JavaScript allows to enter one or several keywords which are then looked up in the index. The resulting matches are presented as IIIF thumbnails, together with the manifest URL and the label metadata. A random selection of keywords is also present.

I also note that this is a completely serverless application, which hosts the necessary JSON statically and reads them into the browser memory upon loading the page. Obviously, this approach has its limitations, but, as with my [Ulpia](https://github.com/TomDeneire/ulpia) project, the benefits of not having to spin up a server for this tool outweigh the disadvantages.

## Remarks

- Not only do institutions seem to neglect IIIF discovery somewhat, some also seem to actively oppose or limit it. I'm not going to name names, but I contacted one institution where the Presentation API was set up with a pager (each collection page offers a link to the next page), but it returned an invalid page at some point. The institution replied that this limit was introduced done on purpose. Sigh.

- Parsing IIIF manifests (both version 2 and 3 manifests are current) with Go has taught me that a lot of institutions seem to implement their own interpretation of the API rather than follow the specifications. Mandatory fields are left out, fields have different data formats (strings instead of arrays and such), and so on.

- I did some experiments with SQLite as a database backend for this application and for the requesting/indexing phase. The first, inspired by the recent [sqlite3 WASM/JS](https://sqlite.org/wasm/doc/ckout/index.md) functionality, I just could not get up and running. The second, I found out, is not a viable option. Even if you insert data into SQLite concurrently with Go routines, SQLite apparently forces everything to sequential writing? Not sure about [this info](https://www.sqlite.org/threadsafe.html), though...

## Wild plan and call to action

Finally, some daydreaming. I made the discovery tool for a sample of the manifests I have collected, but what I would really like to do is push the limits and see how many manifests I can process and still host the index on a static webpage. Currently, for ca. 30,000 manifests, the JSON files are only slightly above 10 MB, so this could definitely be scaled up.

So if you or your instution want to participate in this experiment, or simply deposit your IIIF manifests in the central repository, please get in touch with me.

(I have also contact the people of the Internet Archive to see whether the list of [9.3M IIIF manifests](https://blog.archive.org/2015/10/23/zoom-in-to-9-3-million-internet-archive-books-and-images-through-iiif/) they used to offer, is still available somewhere, but no luck for now!).
