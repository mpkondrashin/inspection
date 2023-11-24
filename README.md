# Inspection

Control Trend Micro Cloud One Network Security Hosted Infrastructure inspection

For network issues diagnostic, it is very often needed to turn off IPS inspection. For Trend Micro Cloud One Network Security Hosted Infrastructure, this option is not available on the management console. Inspection offers this ability as a command line utility.

## Usage 
1. Download [the latest release](https://github.com/mpkondrashin/inspection/releases/latest) of ```inspection``` executable for your platform.
2. Copy ```config_example.yaml``` to ```config.yaml``` in the same directory as Inspection executable itself. Edit ```config.yaml``` and change fields to correct values. 
3. Run ```inspection``` executable with one of following commands: status, on, off.

## Configuration
Inspection offers the following ways to provide configuration parameters:
1. Configuration file ```config.yaml```. Application seeks for this file in its current folder or folder of executable
2. Environment variables
3. Command line parameters

Full config file explained:
```yaml
api_key:    # CloudOne API key (On CloudOne console go to Administration->API Keys->New)
region:     # CloudOne region (On CloudOne console go to Administration-Account Settings->Region)
account_id: # CloudOne account ID (On CloudOne console go to Administration-Account Settings->ID)
aws_region: # AWS region, i.e. us-east-1
```

To set these parameters through the command line, for example, to set AWS Region, use the following command line option:
```commandline 
inspection status --aws_region=us-east-1
```

To set these parameters through the environment variable, add NS prefix. Example for the API Key:
```commandline
NS_API_KEY=tmc12YddE43ASdreseZYhJ5jWAWgaHwBn:5NosR4ed4sdRwe4wfgTYerpedqexms3D14XdqAd8Q5vjcc62irGPHG2weWnh
```

## Output explained

### Turn Inspection Off
```commandline
$./inspection off
2023/11/24 21:35:29 Command: off
2023/11/24 21:35:30 Done
```
### Turn Inspection On
```commandline
$./inspection on
2023/11/24 21:35:29 Command: on
2023/11/24 21:35:30 Done
```

### Status Command

#### Inspection is on

```commandline
$./inspection status
2023/11/24 21:32:18 Command: status
2023/11/24 21:32:19 Action: inspect
2023/11/24 21:32:19 Status: success
2023/11/24 21:32:19 Last change: 2023-11-24 21:09:50 +0200 IST
2023/11/24 21:32:19 Done
```

#### Inspection is off

```commandline
$./inspection status
2023/11/24 21:36:41 Command: status
2023/11/24 21:36:42 Action: bypass
2023/11/24 21:36:42 Status: success
2023/11/24 21:36:42 Last change: 2023-11-24 21:35:31 +0200 IST
2023/11/24 21:36:42 Done
```

## Errors

| Unauthorized                                                | Wrong CloudOne API key or wrong CloudOne region                             |
|-------------------------------------------------------------|-----------------------------------------------------------------------------|
| This request is invalid                                     | Wrong format for CloudOne account or AWS Region                             |
| Forbidden                                                   | Wrong CloudOne account                                                      |
| Data not found                                              | Wrong AWS Region                                                            |
| HTTP request: Get "...": dial tcp: lookup ...: no such host | Nonexistent region or missing Internet connection                           |
| Missing <option name>                                       | Given option is missing in config.yaml, environment, commandline            |
|          