.. Autogenerated by Gollum RST generator (docs/generator/*.go)

JSONToArray
===========

JSONToArray "flattens" a JSON object by selecting specific fields and putting
the values of them into a separated list.

An json input of `{"foo":"value1","bar":"value2"}` can be transformed in a list like `value1,value2`.




Parameters
----------

**Fields**

  The list of all keys which used to create the final text list.
  
  

**Separator** (default: ,)

  This value used as separator for the final text list.
  By default this parameter is set to ",".
  
  

Parameters (from SimpleFormatter)
---------------------------------

**ApplyTo**

  This value chooses the part of the message the formatting should be
  applied to. Use "" to target the message payload; other values specify the name of a metadata field to target.
  By default this parameter is set to "".
  
  

Examples
--------

.. code-block:: yaml

	This example get the `foo` and `bar` fields from a json document
	and create a payload of `foo_value:bar_value`:
	
	 exampleConsumer:
	   Type: consumer.Console
	   Streams: "*"
	   Modulators:
	     - format.JSONToArray
	         Fields:
	           - foo
	           - bar
	         Separator: ;
	
	


