# Terraform Evaluator

An incredibly naive interpreter for evaluating expressions in Terraform.

## How To Use

`docker run webdevwilson/terraform-eval "length([1,2,3])"`

Or, use an alias:

`alias terraform-eval=docker run webdevwilson/terraform-eval "length([1,2,3])"`

## A Few Examples

Try the following commands to test it.

```
> terraform-eval "length([1,2,3])"
3

> terraform-eval "cidrsubnet(\"172.16.0.0/12\", 4, 2)"
172.18.0.0/16

> terraform-eval "formatdate(\"DD MMM YYYY hh:mm ZZZ\", \"2018-01-02T23:12:01Z\")"
02 Jan 2018 23:12 UTC

> terraform-eval "regex(\"^(?:(?P<scheme>[^:/?#]+):)?(?://(?P<authority>[^/?#]*))?\", \"https://terraform.io/docs/\")"
{
  "authority" = "terraform.io"
  "scheme" = "https"
}
```