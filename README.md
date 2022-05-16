# Software Engineer

Hello, and welcome from the Syniti Engineering Team!

Thank you for your interest in joining Syniti, we're working on some
interesting and challenging problems and are looking for engineers with a
growth mindset to join our team. Our interview process helps us learn about
you as a person and at a technical level. But it's just as important that you
learn about us, our work, and our expectations so you know what you're
getting into. We foster a work environment where we are all comfortable and
love coming to work each day. We don't hire lightly and expect everyone to
contribute on day one.

Let's get started!

Please clone or download this repo and use it for the exercises described below.
Once completed you may either send us the URL to your git repo or a zip of
the folder.

We have four questions that we like to ask to get a sense of who you are,
followed by a coding question we use to understand your technical strengths.
You can write your responses to the first four questions directly in the README
or as a separate file.

In answering the code question, please submit code as if you intended to
ship it to production. The details matter. Tests are expected, as is
well-written, simple, idiomatic code. While there are many libraries you
could use, we're expecting to mostly see code that you write yourself. Please
only use critical libraries for common functionality, such as parsing JSON or
writing tests.

We'd recommend you use whatever language you feel strongest in. It doesn't have
to be one we use (mostly Go and JavaScript) — we believe good engineers can be
productive in any language.

Here are the questions, good luck!

1. What’s your proudest achievement? It can be a personal project or something
   you’ve worked on professionally. Just a short paragraph is fine, but I’d
   love to know why you’re proud of it.

- I don't think I have any singularly defined proudest moment in my life.  If
 I did, I don't think it would be associated with work… while I take pride in
 my work and enjoy my accomplishments I live a life outside of work and try not
 to let it define me as a person. I've ran many marathons and placed first sheet,
 I moved across the world on a whim and flourished, etc. I have many achievements
 but none bring me noticeably more pride than any other.

2. What's a personal project you're currently working on? This could be a
   coding side project, hobby, or otherwise real world project you're working
   on.

- I’m actively trying to get back into good marathon shape.

3. Tell us about a technical book or article you read recently, why you liked
   it, and why we should read it.

- 2021 State of DevOps Report. These reports (2014-2021) detail the importance of
 a well built SRE culture and infrastructure at a corporation to allow for the most
 optimal build environment for development. They use studies and data to define 
 best working practices to bolster an optimal development system through empirical
 data. The information presented can help to shape an optimal setup for any software
 company.

4. Tell us about one of your favorite products (physical or software) and one
   specific aspect that makes it truly great.

- I have an ember mug... I enjoy tea and coffee and it maintains a static temperature
 for extended periods of time.  It's a complete luxury item however the one I use
 and appreciate the most on a day to day timeframe. I can forget my tea for an hour
 and come back to it and still have it be piping hot and ready to consume.

5. In this repo is a `data.json` file. It contains an imaginary example set
   of data a customer might need to migrate from one system to another. It's a
   JSON encoded array of objects. The customer understands some of the data
   might be bad and wants to know which records are invalid so they can ensure
   the new system will only have valid data. Write a program that will read
   in the data and mark any records:

   1. That are a duplicate of another record
   2. `name` field is null, missing, or blank
   3. `address` field is null, missing, or blank
   4. `zip` is null, missing, or an invalid U.S. zipcode

   Each record has an ID but that should only be used to identify a record,
   not for validity or duplication testing (eg, two records may be identical
   but have different IDs).

The output of the program should list the IDs of each invalid or duplicate
record, one per line. In the case of duplicates, mark all duplicate IDs.

Example:

```
123ba
439a2
99abc
bac34
```

If you have any questions about the coding questions, please let us know.

Assuming:
- name, address, zip all need to be the same in any singular record for 
it to be a duplicate of another with the same data (ID excluded)
- IDs are always valid and present (nothing explicitly states otherwise
 so no validation done on that field)
- Unexpected errors resulting from bad file load or JSON unmarshalling,
 etc. will also populate the console as an expected error handling measure
 (though not explicitly instructed)
