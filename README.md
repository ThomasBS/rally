Rally API
=========

Rally API is a [Go](http://golang.org/) wrapper for the agile project management [Rally](http://rally1.rallydev.com).

**Warning:** Rally API is under active development!

Features
--------

* Get Rally object
* Fetch references and query references

Requirements
------------

A token for the Rally API must be obtained via [Rally](https://rally1.rallydev.com/login/). Access to Rally is required.


This wrapper does not support basic auth.

Installation
------------

Install Rally API using the ``go get`` command:

    go get github.com/ofmadsen/rally

Usage
-----

Getting a Rally object:

    r := rally.New("rally_obtained_token")
    
    var hr rally.HierarchicalRequirement
    r.Get(&hr, "hierarchical_requirement_id")

All methods returns an ``Error`` on failure or ``nil`` on success.

The objects are not fully populated with data, but contain references and query references that can be called if the relations of the object is required. A reference denotes a one-to-one relation, whereas a query reference denotes as one-to-many relation. Examples of the usage of references:

    // hr is a fetched HierarchicalRequirement object

    // Reference example
    var b rally.Blocker
    r.Fetch(&b, hr.BlockerReference)

    // QueryReference example
    var t rally.TaskQuery
    r.QueryFetch(&t, hr.TasksQueryReference)

Contributing
------------

All is welcome to contribute to this project. Any improvements to the code are appreciated since I'm still a Go rookie.

This project is still under active development and any new features might already be partly finished - please message me before to coordinate.

License
-------

Rally API is available under the [MIT License](http://opensource.org/licenses/MIT).
