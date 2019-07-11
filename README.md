# jwecli
jwe generator

# Motivation
I wanted to know the size of jwe as part of a technical survey.

However, jwe does not have useful services like https://jwt.io.

It was necessary to write a program for generation and it was very lazy.

I thought it would be easier if I put it in a cli form

## Install
`$ go get github.com/convto/jwecli`

## Usage
Usage of jwecli:
```
  -ca string
    	Shorthand content-alg
  -content-alg string
    	Content encryption algorithm. supports:
    	  A128CBC_HS256
  -h	Shorthand help
  -help
    	Print help
  -ka string
    	Shorthand key-alg
  -key-alg string
    	Key encryption algorithm. supports:
    	  RSA1_5
    	  RSA_OAEP
    	  RSA_OAEP_256
    	  ECDH_ES_A128KW
    	  ECDH_ES_A102KW
    	  ECDH_ES_A256KW
  -p string
    	Shorthand payload
  -payload string
    	Payload
```

## Supports algrithms
### content-alg
- A128CBC_HS256
### key-alg
- RSA1_5
- RSA_OAEP
- RSA_OAEP_256
- ECDH_ES_A128KW
- ECDH_ES_A102KW
- ECDH_ES_A256KW

## Example
```bash
# example jwt from https://jwt.io/
$ jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c

# genarate jwe
$ jwecli -key-alg RSA1_5 -content-alg A128CBC_HS256 -payload ${jwt}
eyJhbGciOiJSU0ExXzUiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2In0.0Hk5LPjBCa9xasCo3K30abdhQC_uY8b2yfF7qgHWY55-gIJAvxy4cKHgSnl_TxZwYZiGp4JTvOEYFuaI4UnXtriXN8R1s5VMEN6D73FmPW2QZROHrIcKkfzsRz-FBknrCr6ZoaQN6SCBgw4riYgZych1KJOW-Cojann_11WzjKQioaqYH51cfzTDg6110Qg9aSJDhzYoJmVDm-leich4KOjB-Y-nULc8Z6v-A_ePwKJ1Cmoqs1tOBx83pEsXlMX-RP71YeMUyGefe8-quSYgwycws2YdvFgi-GPrO0N1Y_7x786bHIlfpoDUDXVOfqcKSvAubmAOefXnfJm9dYbJSw.Z-HUevPbrzimWzY1NCK8tA.2jZdvErgElJbH_qidtvx7glOMcuY0_wBLf402ekLTKjM-hqawADkfMeI0_IJi5W2HdBwl5-yG2B14PCUUxF-E3QtY3TjpHrOsLUkLhFDpwVESha_8odN9NOvjnonogoag5hILpjNFY9LSeM6BVOre1gS8bsTnhEVCsatTLc5VW1d4gt9HIkpqF--r9shZXSQiPc8kVeaqm0VlnmBIpqCMQ.pLuwk5LAYQ_h0c1d6XONiw

# check jwe size
$ jwecli -key-alg RSA1_5 -content-alg A128CBC_HS256 -payload ${jwt} | wc -c
     656
```
