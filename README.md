# jmerge

Simple utility for taking in multiple junit xml files and merging them into a
single xml.

## Instal

`go install https://github.com/brentahughes/jmerge@latest`

## Usage

### Standard

`jmerge ./one.xml ./two.xml > merged.xml`

or

`jmerge ./*.xml  > merged.xml`
