.. Autogenerated by Gollum RST generator (docs/generator/*.go)

TextToJSON
==========

This formatter uses a state machine to parse arbitrary text data and
transform it to JSON.




Parameters
----------

**Directive actions**

  Actions are used to write  text read since the last
  transition to the JSON object.
  
  

**Directive flags**

  Flags can modify the parser behavior and can be used to
  store values on a stack accross multiple directives.
  
  

**Directive rules**

  There are some special cases which will cause the parser
  to do additional actions.
  - When writing a value without a key, the state name will become the key.
  - If two keys are written in a row the first key will hold a null value.
  - Writing a key while writing array elements will close the array.
  
  

**Directives**

  Defines an array of directives used to parse text data.
  Each entry must be of the format: "State:Token:NextState:Flags:Function".
  State denotes the name of the state owning this entry. Multiple entries per
  state are allowed. Token holds a string that triggers a state transition.
  NextState holds the target of the state transition. Flags is an optional
  field and is used to trigger special parser behavior. Flags can be comma
  separated if you need to use more than one.
  Function defines an action that is triggered upon state transition.
  Spaces will be stripped from all fields but Token. If a fields requires a
  colon it has to be escaped with a backslash. Other escape characters
  supported are \n, \r and \t.
  By default this parameter is set to an empty list.
  
  

**StartState**

  Defines the name of the initial state when parsing a message.
  When set to an empty string the first state from the directives array will
  be used.
  By default this parameter is set to "".
  
  

**TimestampRead**

  Defines a time.Parse compatible format string used to read
  time fields when using the "dat" directive.
  By default this parameter is set to "20060102150405".
  
  

**TimestampWrite** (default: 2006-01-02 15:04:05 MST)

  Defines a time.Format compatible format string used to
  write time fields when using the "dat" directive.
  By default this parameter is set to "2006-01-02 15:04:05 MST".
  
  

**UnixTimestampRead**

  Defines the unix timestamp format expected from fields
  that are parsed using the "dat" directive. Valid valies are "s" for seconds,
  "ms" for milliseconds, or "ns" for nanoseconds. This parameter is ignored
  unless TimestampRead is set to "".
  By default this parameter is set to "".
  
  

**append**

  Append the token to the current match and continue reading.
  
  

**arr**

  Start a new array.
  
  

**arr+dat**

  arr followed by dat.
  
  

**arr+esc**

  arr followed by esc.
  
  

**arr+val**

  arr followed by val.
  
  

**continue**

  Prepend the token to the next match.
  
  

**dat**

  Write the parsed section as a timestamp value.
  
  

**dat+end**

  dat followed by end.
  
  

**end**

  Close an array or object.
  
  

**esc**

  Write the parsed section as a escaped string value.
  
  

**esc+end**

  esc followed by end.
  
  

**include**

  Append the token to the current match.
  
  

**key**

  Write the parsed section as a key.
  
  

**obj**

  Start a new object.
  
  

**pop**

  Pop the stack and use the returned state if possible.
  
  

**push**

  Push the current state to the stack.
  
  

**val**

  Write the parsed section as a value without quotes.
  
  

**val+end**

  val followed by end.
  
  

Parameters (from SimpleFormatter)
---------------------------------

**ApplyTo**

  This value chooses the part of the message the formatting should be
  applied to. Use "" to target the message payload; other values specify the name of a metadata field to target.
  By default this parameter is set to "".
  
  

Examples
--------

.. code-block:: yaml

	The following example parses JSON data.
	
	 ExampleConsumer:
	   Type: consumer.Console
	   Streams: console
	   Modulators:
	     - format.JSON:
	       Directives:
	         - "findKey   :\":  key       :      :        "
	         - "findKey   :}:             : pop  : end    "
	         - "key       :\":  findVal   :      : key    "
	         - "findVal   :\\:: value     :      :        "
	         - "value     :\":  string    :      :        "
	         - "value     :[:   array     : push : arr    "
	         - "value     :{:   findKey   : push : obj    "
	         - "value     :,:   findKey   :      : val    "
	         - "value     :}:             : pop  : val+end"
	         - "string    :\":  findKey   :      : esc    "
	         - "array     :[:   array     : push : arr    "
	         - "array     :{:   findKey   : push : obj    "
	         - "array     :]:             : pop  : val+end"
	         - "array     :,:   array     :      : val    "
	         - "array     :\":  arrString :      :        "
	         - "arrString :\":  array     :      : esc    "
	
	


