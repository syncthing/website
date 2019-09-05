---
title: New GUI & REST API
header: true
date: 2019-09-05
state: Accepting Applications until 2019-10-05
---

Syncthing uses a web based **GUI** (currently HTML, Bootstrap 3, JavaScript,
and Angular 1.3). This GUI is a bit old, has grown "organically" (that is,
without planning or aforethought) for several years, and has multiple
issues.

The underlying **REST API** has some similar issues, in that it also has
grown from something small and simple to something large and complex without
the benefit of planning.

We would now like to fix these issues with a redesign of both, in concert.

<!--more-->

Specifically, we are looking for:

- A **GUI design refresh**, taking into account the features that exist
  today and some limitations that should not persist into the new version.
  One such limitation is that performance is currently very bad for setups
  with many devices and/or folders. The refresh should bring both a visually
  more delightful look and a more well considered organization and some new
  features that are sorely lacking today, for example a getting-started
  "wizard" or guide. Some ideas already exist from which to draw
  inspiration, but none are 100% complete and we appreciate you taking
  initiative on this point. The GUI should follow modern accessibility
  standards and be responsive enough to work well on both mobile and
  desktop.

- A **REST API design refresh**, to better support the new GUI but also
  keeping in mind the needs of external integrations, including other
  potential GUI implementations. The API should provide easy access to
  current internal state in the form of events, conveniently allow common
  actions and configuration changes, and follow a structured plan. The
  overall architecture remains based on HTTP verbs
  (GET/POST/PUT/DELETE/PATCH) and uses JSON serialization.

- An **implementation of the GUI refresh** in a suitable, modern,
  environment. We still need to run in a browser, but we are open to
  suggestions on the front end technologies used (React? Vue.js? Angular?
  Dart? etc). We expect the implementation to follow best current practices
  in the chosen framework, whichever it ends up being. Providing a
  suggestion on this point and a couple of references of previous work is
  meriting.

- In concert with the GUI refresh, an **implementation of the REST API
  refresh**. This work is done in the existing Syncthing backend, in Go. We
  expect the GUI and new REST API to evolve together over the course of the
  project.

- The work to proceed using the established **open source processes** in use
  by the Syncthing project. This means Git, GitHub, and pull requests.
  Previous open source experience is meriting. We will regardless assist all
  along the way, both with procedure and code reviews.

To be clear, this is a from-scratch rewrite of the GUI, not a careful
piece-by-piece improvement of the existing code. There will be new, exciting
bugs and we will need to provide the user with the option to use either the
new or the old GUI for quite a while. The new REST API will similarly live
beside the old one, for compatibility reasons, for some time.

The GUI and REST API design will be iterated in cooperation with the project
maintainers and a more detailed implementation plan created as part of the
project. We will also assist with guiding, mentoring, testing and -- as
required -- the implementation of backend API changes.

## Compensation

This is a time-limited grant, not an indefinite employment. We expect the
project to require several months, up to maximum about half a year. We will
provide a salary or consultancy rate in line with market rates for a junior
to mid-level developer. Correspondingly we will factor in time for learning,
experimentation and mentoring of a more junior developer.

## Eligibility & Application

The grant is available for individuals only, corporations and similar
constructions are not eligible. Individuals serving on the board of the
Syncthing Foundation or employed by Kastelo Inc. are _not_ eligible.

Submit your application via email to
[foundation@syncthing.org](mailto:foundation@syncthing.org). Please include
at minimum:

- Basic information about yourself and your background,
- Suggestion of front end technologies you would use in this project,
- Any relevant references to prior work in the area,
- Your salary or compensation rate requirements.

If you have questions not answered by this brief please see [this forum
thread](https://forum.syncthing.net/t/syncthing-foundation-grant-new-gui-rest-api/13732)
where the grant is discussed, or feel free to use the contact email above if
you want to ask a question privately.

The evaluation process will include thorough evaluation of the written
application as well as personal interviews and consultations with external
international reviewers. Evaluation panel assessment will be in complete
confidentiality, and will not be disclosed. The decision on the accepted
grant application will be taken by the Board of the Syncthing Foundation.

## Sponsors

This grant is sponsored financially by [Kastelo Inc](https://kastelo.net/).
Kastelo provides consulting and support services for, among other things,
Syncthing.
