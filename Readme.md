<!--
 Copyright 2020 Kaan Karakaya
 
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
 
     http://www.apache.org/licenses/LICENSE-2.0
 
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

# Gkey

Gkey is a simple vaultless password manager in Go.
It's generating sha256 sum of password + "-" + realm combination.
And with that sum gkey generating secure password.

## Usage

### Example

```bash
gkey -r github.com -p my-super-secure-master-password

# 6a?_#ZH$kVtB,P*!
```

### Flags

`-p <master password>` - master password which will be used to generate other passwords

`-r <realm>` - any string which identifies requested password, most likely key usage or resource URL

`-l <length>` - number of characters in the generated password (default 16 )
