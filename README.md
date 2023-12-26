
gofritz
===========================

A Go library, to interact with AVM Fritz!Box routers/devices via TR064 API.
The library offers convenient Go-functions, which are SOAP based client stubs,
having no external dependencies, and purely work with Go's standard lib.
In future releases, a command line (CLI) implementation is planned.

Contributors and testers are highly welcome.

This work was inspired by https://github.com/kanimaru/fritzbox-soap-example - thanks for that.

#### Implementation Status

* Tested with `Fritz!Box Cable` and a `Fritz!Box` device
* The CLI part is not yet implemented
* The library part is fully implemented (using code generation, based on AVM's service descriptions)
* The error handling is rather naive, and mostly creates a panic (TODO: for future releases, the errors should rather percolate up to the lib consumers)
* There are no config options for timeouts, or retries or other things yet, to make the lib more robust
* No anonymous API calls supported (only with user/pass)
* No SSL transport encryption is supported (planned for future releases)

#### Example usage

The typical pattern of how to use this library is to create a session object first, and then use it for multiple API calls.
The `session` object does contain username/password to automatically authenticate against the API.
Because the authentication schema implemented by Fritz!Box is based on a challenge response method with a derived session,
it's recommended to re-use the session object for multiple API calls, to avoid redundant login API calls.

```go
 username := os.Getenv("FB_USERNAME")
 password := os.Getenv("FB_PASSWORD")
 session := soap.NewSession("fritz.box", username, password)
 
 hostList, err := lan.XAvmGetHostList(session)
 if err != nil {
     panic(err)
 }
 for _, host := range hostList {
     println(fmt.Sprintf("Host: %s, Active:%v", host.XAvmPort, host.Active))
 }
```

## Configure the TR064 API and a user

This is the basic setup to be done, so an application can use this API.

1. log into http://fritz.box/ via your admin account
2. enable TR064, via menu path... 
   'home network' -> 'network settings' -> 'advanced settings' -> tick the box 'allow access for applications'
3. create a user with password, via menu path...
   'system' -> 'Fritz!box user' -> edit or create a new one and set a password

## Supported API functions

This is the list of available functions, which are available via TR064 API.
Additionally, some comfort-functions exists, ease the usage of the API (see below).

The column 'discovered via' indicates, since when the API is available.
Some APIs are also specific for the 'cable' model and not available by other models.
For future releases its planned to enhance the documentation, for which model a function is available.
For now, please consult the Go docs of the implementation, which links the .XML file function description,
to verify your device supports the SOAP call.

<!-- MARKDOWN-AUTO-DOCS:START (CODE:src=functions_registry.go) -->
| (Go) function | Discovered via |
|---------------|----------------|
| [gateway.GetDeviceInfo](./pkg/tr064/gateway/get_device_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetProvisioningCode](./pkg/tr064/gateway/set_provisioning_code.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetDeviceLog](./pkg/tr064/gateway/get_device_log.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetSecurityPort](./pkg/tr064/gateway/get_security_port.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetPersistentData](./pkg/tr064/gateway/get_persistent_data.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetPersistentData](./pkg/tr064/gateway/set_persistent_data.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.ConfigurationStarted](./pkg/tr064/gateway/configuration_started.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.ConfigurationFinished](./pkg/tr064/gateway/configuration_finished.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.FactoryReset](./pkg/tr064/gateway/factory_reset.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.Reboot](./pkg/tr064/gateway/reboot.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.X_GenerateUUID](./pkg/tr064/gateway/x_generate_uuid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetConfigFile](./pkg/tr064/gateway/xavm_get_config_file.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmSetConfigFile](./pkg/tr064/gateway/xavm_set_config_file.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmCreateUrlSID](./pkg/tr064/gateway/xavm_create_url_sid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmSendSupportData](./pkg/tr064/gateway/xavm_send_support_data.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetSupportDataInfo](./pkg/tr064/gateway/xavm_get_support_data_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetSupportDataEnable](./pkg/tr064/gateway/xavm_get_support_data_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmSetSupportDataEnable](./pkg/tr064/gateway/xavm_set_support_data_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetDefaultConnectionService](./pkg/tr064/gateway/set_default_connection_service.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetDefaultConnectionService](./pkg/tr064/gateway/get_default_connection_service.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetForwardNumberOfEntries](./pkg/tr064/gateway/get_forward_number_of_entries.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.AddForwardingEntry](./pkg/tr064/gateway/add_forwarding_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.DeleteForwardingEntry](./pkg/tr064/gateway/delete_forwarding_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetSpecificForwardingEntry](./pkg/tr064/gateway/get_specific_forwarding_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetGenericForwardingEntry](./pkg/tr064/gateway/get_generic_forwarding_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetForwardingEntryEnable](./pkg/tr064/gateway/set_forwarding_entry_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetLanConfigSecurityInfo](./pkg/tr064/gateway/get_lan_config_security_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetCurrentUser](./pkg/tr064/gateway/xavm_get_current_user.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetAnonymousLogin](./pkg/tr064/gateway/xavm_get_anonymous_login.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetConfigPassword](./pkg/tr064/gateway/set_config_password.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetUserList](./pkg/tr064/gateway/xavm_get_user_list.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetManagementServerInfo](./pkg/tr064/gateway/get_management_server_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetManagementServerURL](./pkg/tr064/gateway/set_management_server_url.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetManagementServerUsername](./pkg/tr064/gateway/set_management_server_username.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetManagementServerPassword](./pkg/tr064/gateway/set_management_server_password.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetPeriodicInform](./pkg/tr064/gateway/set_periodic_inform.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetConnectionRequestAuthentication](./pkg/tr064/gateway/set_connection_request_authentication.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetUpgradeManagement](./pkg/tr064/gateway/set_upgrade_management.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.X_SetTR069Enable](./pkg/tr064/gateway/x_set_tr069_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetTR069FirmwareDownloadEnabled](./pkg/tr064/gateway/xavm_get_tr069_firmware_download_enabled.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmSetTR069FirmwareDownloadEnabled](./pkg/tr064/gateway/xavm_set_tr069_firmware_download_enabled.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetTimeInfo](./pkg/tr064/gateway/get_time_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetNTPServers](./pkg/tr064/gateway/set_ntp_servers.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetUserInterfaceInfo](./pkg/tr064/gateway/get_user_interface_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmCheckUpdate](./pkg/tr064/gateway/xavm_check_update.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmDoUpdate](./pkg/tr064/gateway/xavm_do_update.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmDoPrepareCGI](./pkg/tr064/gateway/xavm_do_prepare_c_g_i.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmDoManualUpdate](./pkg/tr064/gateway/xavm_do_manual_update.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetInternationalConfig](./pkg/tr064/gateway/xavm_get_international_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmSetInternationalConfig](./pkg/tr064/gateway/xavm_set_international_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetUserInterfaceInfo](./pkg/tr064/gateway/xavm_get_user_interface_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmSetConfig](./pkg/tr064/gateway/xavm_set_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAvmStorageInfo](./pkg/tr064/gateway/get_avm_storage_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.RequestFTPServerWAN](./pkg/tr064/gateway/request_ftp_server_wan.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetFTPServer](./pkg/tr064/gateway/set_ftp_server.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetFTPServerWAN](./pkg/tr064/gateway/set_ftp_server_wan.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetSMBServer](./pkg/tr064/gateway/set_smb_server.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetUserInfo](./pkg/tr064/gateway/get_user_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetUserConfig](./pkg/tr064/gateway/set_user_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAvmWebdavInfo](./pkg/tr064/gateway/get_avm_webdav_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetAvmWebdavConfig](./pkg/tr064/gateway/set_avm_webdav_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAvmUpnpInfo](./pkg/tr064/gateway/get_avm_upnp_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetAvmUpnpConfig](./pkg/tr064/gateway/set_avm_upnp_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAvmSpeedtestInfo](./pkg/tr064/gateway/get_avm_speedtest_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetAvmSpeedtestConfig](./pkg/tr064/gateway/set_avm_speedtest_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAvmSpeedtestStatistics](./pkg/tr064/gateway/get_avm_speedtest_statistics.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.ResetStatistics](./pkg/tr064/gateway/reset_statistics.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAvmRemoteAccessInfo](./pkg/tr064/gateway/get_avm_remote_access_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetAvmRemoteAccessConfig](./pkg/tr064/gateway/set_avm_remote_access_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetAvmRemoteAccessEnable](./pkg/tr064/gateway/set_avm_remote_access_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetLetsEncryptEnable](./pkg/tr064/gateway/set_lets_encrypt_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetDDNSInfo](./pkg/tr064/gateway/get_d_dns_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetDDNSProviders](./pkg/tr064/gateway/get_d_dns_providers.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetDDNSConfig](./pkg/tr064/gateway/set_d_dns_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAvmMyFritzInfo](./pkg/tr064/gateway/get_avm_my_fritz_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetMyFRITZ](./pkg/tr064/gateway/set_my_fritz.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetNumberOfServices](./pkg/tr064/gateway/get_number_of_services.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetServiceByIndex](./pkg/tr064/gateway/get_service_by_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetServiceByIndex](./pkg/tr064/gateway/set_service_by_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.DeleteServiceByIndex](./pkg/tr064/gateway/delete_service_by_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetInfoEx](./pkg/tr064/gateway/get_info_ex.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmAddVoIPAccount](./pkg/tr064/gateway/xavm_add_voip_account.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetVoIPAccount](./pkg/tr064/gateway/xavm_get_voip_account.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmDelVoIPAccount](./pkg/tr064/gateway/xavm_del_voip_account.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetVoIPAccounts](./pkg/tr064/gateway/xavm_get_voip_accounts.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetVoIPStatus](./pkg/tr064/gateway/xavm_get_voip_status.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetXVoipInfo](./pkg/tr064/gateway/get_x_voip_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetXVoipConfig](./pkg/tr064/gateway/set_x_voip_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetMaxVoIPNumbers](./pkg/tr064/gateway/get_max_voip_numbers.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetExistingVoIPNumbers](./pkg/tr064/gateway/get_existing_voip_numbers.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetNumberOfClients](./pkg/tr064/gateway/xavm_get_number_of_clients.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetClient](./pkg/tr064/gateway/xavm_get_client.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetClient2](./pkg/tr064/gateway/xavm_get_client2.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmSetClient](./pkg/tr064/gateway/xavm_set_client.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmSetClient2](./pkg/tr064/gateway/xavm_set_client2.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetClient3](./pkg/tr064/gateway/xavm_get_client3.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetClientByClientId](./pkg/tr064/gateway/xavm_get_client_by_client_id.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmSetClient3](./pkg/tr064/gateway/xavm_set_client3.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmSetClient4](./pkg/tr064/gateway/xavm_set_client4.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetClients](./pkg/tr064/gateway/xavm_get_clients.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetNumberOfNumbers](./pkg/tr064/gateway/xavm_get_number_of_numbers.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetNumbers](./pkg/tr064/gateway/xavm_get_numbers.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmDeleteClient](./pkg/tr064/gateway/xavm_delete_client.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmDialGetConfig](./pkg/tr064/gateway/xavm_dial_get_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmDialSetConfig](./pkg/tr064/gateway/xavm_dial_set_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmDialNumber](./pkg/tr064/gateway/xavm_dial_number.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmDialHangup](./pkg/tr064/gateway/xavm_dial_hangup.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetPhonePort](./pkg/tr064/gateway/xavm_get_phone_port.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmSetDelayedCallNotification](./pkg/tr064/gateway/xavm_set_delayed_call_notification.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetVoIPCommonCountryCode](./pkg/tr064/gateway/get_voip_common_country_code.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetVoIPCommonCountryCode](./pkg/tr064/gateway/xavm_get_voip_common_country_code.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetVoIPCommonCountryCode](./pkg/tr064/gateway/set_voip_common_country_code.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmSetVoIPCommonCountryCode](./pkg/tr064/gateway/xavm_set_voip_common_country_code.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetVoIPEnableCountryCode](./pkg/tr064/gateway/get_voip_enable_country_code.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetVoIPEnableCountryCode](./pkg/tr064/gateway/set_voip_enable_country_code.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetVoIPCommonAreaCode](./pkg/tr064/gateway/get_voip_common_area_code.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetVoIPCommonAreaCode](./pkg/tr064/gateway/xavm_get_voip_common_area_code.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetVoIPCommonAreaCode](./pkg/tr064/gateway/set_voip_common_area_code.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmSetVoIPCommonAreaCode](./pkg/tr064/gateway/xavm_set_voip_common_area_code.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetVoIPEnableAreaCode](./pkg/tr064/gateway/get_voip_enable_area_code.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetVoIPEnableAreaCode](./pkg/tr064/gateway/set_voip_enable_area_code.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetAlarmClock](./pkg/tr064/gateway/xavm_get_alarm_clock.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmSetAlarmClockEnable](./pkg/tr064/gateway/xavm_set_alarm_clock_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.XavmGetNumberOfAlarmClocks](./pkg/tr064/gateway/xavm_get_number_of_alarm_clocks.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAvmOnTelInfo](./pkg/tr064/gateway/get_avm_on_tel_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetAvmOnTelEnable](./pkg/tr064/gateway/set_avm_on_tel_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetAvmOnTelConfig](./pkg/tr064/gateway/set_avm_on_tel_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetInfoByIndex](./pkg/tr064/gateway/get_info_by_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetEnableByIndex](./pkg/tr064/gateway/set_enable_by_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetConfigByIndex](./pkg/tr064/gateway/set_config_by_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.DeleteByIndex](./pkg/tr064/gateway/delete_by_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetNumberOfEntries](./pkg/tr064/gateway/get_number_of_entries.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetCallList](./pkg/tr064/gateway/get_call_list.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetPhonebookList](./pkg/tr064/gateway/get_phonebook_list.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetPhonebook](./pkg/tr064/gateway/get_phonebook.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.AddPhonebook](./pkg/tr064/gateway/add_phonebook.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.DeletePhonebook](./pkg/tr064/gateway/delete_phonebook.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetPhonebookEntry](./pkg/tr064/gateway/get_phonebook_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetPhonebookEntryUID](./pkg/tr064/gateway/get_phonebook_entry_uid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetPhonebookEntry](./pkg/tr064/gateway/set_phonebook_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetPhonebookEntryUID](./pkg/tr064/gateway/set_phonebook_entry_uid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.DeletePhonebookEntry](./pkg/tr064/gateway/delete_phonebook_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.DeletePhonebookEntryUID](./pkg/tr064/gateway/delete_phonebook_entry_uid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetCallBarringEntry](./pkg/tr064/gateway/get_call_barring_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetCallBarringEntryByNum](./pkg/tr064/gateway/get_call_barring_entry_by_num.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetCallBarringList](./pkg/tr064/gateway/get_call_barring_list.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetCallBarringEntry](./pkg/tr064/gateway/set_call_barring_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.DeleteCallBarringEntryUID](./pkg/tr064/gateway/delete_call_barring_entry_uid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetDECTHandsetList](./pkg/tr064/gateway/get_dect_handset_list.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetDECTHandsetInfo](./pkg/tr064/gateway/get_dect_handset_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetDECTHandsetPhonebook](./pkg/tr064/gateway/set_dect_handset_phonebook.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetNumberOfDeflections](./pkg/tr064/gateway/get_number_of_deflections.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetDeflection](./pkg/tr064/gateway/get_deflection.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetDeflections](./pkg/tr064/gateway/get_deflections.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetDeflectionEnable](./pkg/tr064/gateway/set_deflection_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetNumberOfDectEntries](./pkg/tr064/gateway/get_number_of_dect_entries.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetGenericDectEntry](./pkg/tr064/gateway/get_generic_dect_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetSpecificDectEntry](./pkg/tr064/gateway/get_specific_dect_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.DectDoUpdate](./pkg/tr064/gateway/dect_do_update.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetDectListPath](./pkg/tr064/gateway/get_dect_list_path.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAvmTamInfo](./pkg/tr064/gateway/get_avm_tam_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetAvmTamEnable](./pkg/tr064/gateway/set_avm_tam_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetMessageList](./pkg/tr064/gateway/get_message_list.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.MarkMessage](./pkg/tr064/gateway/mark_message.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.DeleteMessage](./pkg/tr064/gateway/delete_message.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetList](./pkg/tr064/gateway/get_list.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAvmAppSetupInfo](./pkg/tr064/gateway/get_avm_app_setup_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetConfig](./pkg/tr064/gateway/get_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAppMessageFilter](./pkg/tr064/gateway/get_app_message_filter.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.RegisterApp](./pkg/tr064/gateway/register_app.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetAppVPN](./pkg/tr064/gateway/set_app_vpn.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetAppVPNwithPFS](./pkg/tr064/gateway/set_app_vpnwith_pfs.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetAppMessageFilter](./pkg/tr064/gateway/set_app_message_filter.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetAppMessageReceiver](./pkg/tr064/gateway/set_app_message_receiver.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.ResetEvent](./pkg/tr064/gateway/reset_event.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAppRemoteInfo](./pkg/tr064/gateway/get_app_remote_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetBoxSenderId](./pkg/tr064/gateway/get_box_sender_id.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAvmHomeautoInfo](./pkg/tr064/gateway/get_avm_homeauto_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetGenericDeviceInfos](./pkg/tr064/gateway/get_generic_device_infos.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetSpecificDeviceInfos](./pkg/tr064/gateway/get_specific_device_infos.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetDeviceName](./pkg/tr064/gateway/set_device_name.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetSwitch](./pkg/tr064/gateway/set_switch.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetNumberOfDeviceEntries](./pkg/tr064/gateway/get_number_of_device_entries.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetGenericDeviceEntry](./pkg/tr064/gateway/get_generic_device_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetSpecificDeviceEntry](./pkg/tr064/gateway/get_specific_device_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.DeviceDoUpdate](./pkg/tr064/gateway/device_do_update.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetNumberOfFilelinkEntries](./pkg/tr064/gateway/get_number_of_filelink_entries.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetGenericFilelinkEntry](./pkg/tr064/gateway/get_generic_filelink_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetSpecificFilelinkEntry](./pkg/tr064/gateway/get_specific_filelink_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.NewFilelinkEntry](./pkg/tr064/gateway/new_filelink_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetFilelinkEntry](./pkg/tr064/gateway/set_filelink_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.DeleteFilelinkEntry](./pkg/tr064/gateway/delete_filelink_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetFilelinkListPath](./pkg/tr064/gateway/get_filelink_list_path.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAvmAuthInfo](./pkg/tr064/gateway/get_avm_auth_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetState](./pkg/tr064/gateway/get_state.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetAvmAuthConfig](./pkg/tr064/gateway/set_avm_auth_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.MarkTicket](./pkg/tr064/gateway/mark_ticket.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetTicketIDStatus](./pkg/tr064/gateway/get_ticket_id_status.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.DiscardAllTickets](./pkg/tr064/gateway/discard_all_tickets.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.DisallowWANAccessByIP](./pkg/tr064/gateway/disallow_wan_access_by_ip.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetWANAccessByIP](./pkg/tr064/gateway/get_wan_access_by_ip.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAvmMediaInfo](./pkg/tr064/gateway/get_avm_media_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetDVBCEnable](./pkg/tr064/gateway/get_dvbc_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetDVBCEnable](./pkg/tr064/gateway/set_dvbc_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.StationSearch](./pkg/tr064/gateway/station_search.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetSearchProgress](./pkg/tr064/gateway/get_search_progress.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetAvmUspControllerInfo](./pkg/tr064/gateway/get_avm_usp_controller_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetUSPControllerByIndex](./pkg/tr064/gateway/get_usp_controller_by_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.GetUSPControllerNumberOfEntries](./pkg/tr064/gateway/get_usp_controller_number_of_entries.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.AddUSPController](./pkg/tr064/gateway/add_usp_controller.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.DeleteUSPControllerByIndex](./pkg/tr064/gateway/delete_usp_controller_by_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [gateway.SetUSPControllerEnableByIndex](./pkg/tr064/gateway/set_usp_controller_enable_by_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1SetWlanConfigurationEnable](./pkg/tr064/lan/wlan1_set_wlan_configuration_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1GetWlanConfigurationInfo](./pkg/tr064/lan/wlan1_get_wlan_configuration_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1SetWlanConfigurationConfig](./pkg/tr064/lan/wlan1_set_wlan_configuration_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1SetSecurityKeys](./pkg/tr064/lan/wlan1_set_security_keys.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1GetSecurityKeys](./pkg/tr064/lan/wlan1_get_security_keys.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1SetDefaultWEPKeyIndex](./pkg/tr064/lan/wlan1_set_default_wep_key_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1GetDefaultWEPKeyIndex](./pkg/tr064/lan/wlan1_get_default_wep_key_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1SetBasBeaconSecurityProperties](./pkg/tr064/lan/wlan1_set_bas_beacon_security_properties.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1GetBasBeaconSecurityProperties](./pkg/tr064/lan/wlan1_get_bas_beacon_security_properties.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1GetWlanConfigurationStatistics](./pkg/tr064/lan/wlan1_get_wlan_configuration_statistics.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1GetPacketStatistics](./pkg/tr064/lan/wlan1_get_packet_statistics.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1GetBSSID](./pkg/tr064/lan/wlan1_get_b_ssid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1GetSSID](./pkg/tr064/lan/wlan1_get_ssid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1SetSSID](./pkg/tr064/lan/wlan1_set_ssid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1GetBeaconType](./pkg/tr064/lan/wlan1_get_beacon_type.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1SetBeaconType](./pkg/tr064/lan/wlan1_set_beacon_type.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1GetChannelInfo](./pkg/tr064/lan/wlan1_get_channel_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1SetChannel](./pkg/tr064/lan/wlan1_set_channel.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1GetBeaconAdvertisement](./pkg/tr064/lan/wlan1_get_beacon_advertisement.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1SetBeaconAdvertisement](./pkg/tr064/lan/wlan1_set_beacon_advertisement.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1GetTotalAssociations](./pkg/tr064/lan/wlan1_get_total_associations.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1GetGenericAssociatedDeviceInfo](./pkg/tr064/lan/wlan1_get_generic_associated_device_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1GetSpecificAssociatedDeviceInfo](./pkg/tr064/lan/wlan1_get_specific_associated_device_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1XavmGetSpecificAssociatedDeviceInfoByIp](./pkg/tr064/lan/wlan1_xavm_get_specific_associated_device_info_by_ip.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1XavmGetWLANDeviceListPath](./pkg/tr064/lan/wlan1_xavm_get_wlan_device_list_path.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1XavmSetStickSurfEnable](./pkg/tr064/lan/wlan1_xavm_set_stick_surf_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1XavmGetIPTVOptimized](./pkg/tr064/lan/wlan1_xavm_get_ip_tv_optimized.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1XavmSetIPTVOptimized](./pkg/tr064/lan/wlan1_xavm_set_ip_tv_optimized.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1XavmGetNightControl](./pkg/tr064/lan/wlan1_xavm_get_night_control.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1XavmGetWLANHybridMode](./pkg/tr064/lan/wlan1_xavm_get_wlan_hybrid_mode.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1XavmSetWLANHybridMode](./pkg/tr064/lan/wlan1_xavm_set_wlan_hybrid_mode.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1XavmGetWLANExtInfo](./pkg/tr064/lan/wlan1_xavm_get_wlan_ext_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1XavmGetWPSInfo](./pkg/tr064/lan/wlan1_xavm_get_wps_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1XavmSetWPSConfig](./pkg/tr064/lan/wlan1_xavm_set_wps_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1XavmSetWPSEnable](./pkg/tr064/lan/wlan1_xavm_set_wps_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1XavmSetWLANGlobalEnable](./pkg/tr064/lan/wlan1_xavm_set_wlan_global_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan1XavmGetWLANConnectionInfo](./pkg/tr064/lan/wlan1_xavm_get_wlan_connection_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2SetWlanConfigurationEnable](./pkg/tr064/lan/wlan2_set_wlan_configuration_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2GetWlanConfigurationInfo](./pkg/tr064/lan/wlan2_get_wlan_configuration_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2SetWlanConfigurationConfig](./pkg/tr064/lan/wlan2_set_wlan_configuration_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2SetSecurityKeys](./pkg/tr064/lan/wlan2_set_security_keys.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2GetSecurityKeys](./pkg/tr064/lan/wlan2_get_security_keys.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2SetDefaultWEPKeyIndex](./pkg/tr064/lan/wlan2_set_default_wep_key_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2GetDefaultWEPKeyIndex](./pkg/tr064/lan/wlan2_get_default_wep_key_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2SetBasBeaconSecurityProperties](./pkg/tr064/lan/wlan2_set_bas_beacon_security_properties.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2GetBasBeaconSecurityProperties](./pkg/tr064/lan/wlan2_get_bas_beacon_security_properties.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2GetWlanConfigurationStatistics](./pkg/tr064/lan/wlan2_get_wlan_configuration_statistics.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2GetPacketStatistics](./pkg/tr064/lan/wlan2_get_packet_statistics.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2GetBSSID](./pkg/tr064/lan/wlan2_get_b_ssid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2GetSSID](./pkg/tr064/lan/wlan2_get_ssid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2SetSSID](./pkg/tr064/lan/wlan2_set_ssid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2GetBeaconType](./pkg/tr064/lan/wlan2_get_beacon_type.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2SetBeaconType](./pkg/tr064/lan/wlan2_set_beacon_type.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2GetChannelInfo](./pkg/tr064/lan/wlan2_get_channel_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2SetChannel](./pkg/tr064/lan/wlan2_set_channel.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2GetBeaconAdvertisement](./pkg/tr064/lan/wlan2_get_beacon_advertisement.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2SetBeaconAdvertisement](./pkg/tr064/lan/wlan2_set_beacon_advertisement.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2GetTotalAssociations](./pkg/tr064/lan/wlan2_get_total_associations.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2GetGenericAssociatedDeviceInfo](./pkg/tr064/lan/wlan2_get_generic_associated_device_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2GetSpecificAssociatedDeviceInfo](./pkg/tr064/lan/wlan2_get_specific_associated_device_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2XavmGetSpecificAssociatedDeviceInfoByIp](./pkg/tr064/lan/wlan2_xavm_get_specific_associated_device_info_by_ip.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2XavmGetWLANDeviceListPath](./pkg/tr064/lan/wlan2_xavm_get_wlan_device_list_path.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2XavmSetStickSurfEnable](./pkg/tr064/lan/wlan2_xavm_set_stick_surf_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2XavmGetIPTVOptimized](./pkg/tr064/lan/wlan2_xavm_get_ip_tv_optimized.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2XavmSetIPTVOptimized](./pkg/tr064/lan/wlan2_xavm_set_ip_tv_optimized.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2XavmGetNightControl](./pkg/tr064/lan/wlan2_xavm_get_night_control.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2XavmGetWLANHybridMode](./pkg/tr064/lan/wlan2_xavm_get_wlan_hybrid_mode.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2XavmSetWLANHybridMode](./pkg/tr064/lan/wlan2_xavm_set_wlan_hybrid_mode.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2XavmGetWLANExtInfo](./pkg/tr064/lan/wlan2_xavm_get_wlan_ext_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2XavmGetWPSInfo](./pkg/tr064/lan/wlan2_xavm_get_wps_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2XavmSetWPSConfig](./pkg/tr064/lan/wlan2_xavm_set_wps_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2XavmSetWPSEnable](./pkg/tr064/lan/wlan2_xavm_set_wps_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2XavmSetWLANGlobalEnable](./pkg/tr064/lan/wlan2_xavm_set_wlan_global_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan2XavmGetWLANConnectionInfo](./pkg/tr064/lan/wlan2_xavm_get_wlan_connection_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3SetWlanConfigurationEnable](./pkg/tr064/lan/wlan3_set_wlan_configuration_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3GetWlanConfigurationInfo](./pkg/tr064/lan/wlan3_get_wlan_configuration_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3SetWlanConfigurationConfig](./pkg/tr064/lan/wlan3_set_wlan_configuration_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3SetSecurityKeys](./pkg/tr064/lan/wlan3_set_security_keys.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3GetSecurityKeys](./pkg/tr064/lan/wlan3_get_security_keys.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3SetDefaultWEPKeyIndex](./pkg/tr064/lan/wlan3_set_default_wep_key_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3GetDefaultWEPKeyIndex](./pkg/tr064/lan/wlan3_get_default_wep_key_index.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3SetBasBeaconSecurityProperties](./pkg/tr064/lan/wlan3_set_bas_beacon_security_properties.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3GetBasBeaconSecurityProperties](./pkg/tr064/lan/wlan3_get_bas_beacon_security_properties.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3GetWlanConfigurationStatistics](./pkg/tr064/lan/wlan3_get_wlan_configuration_statistics.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3GetPacketStatistics](./pkg/tr064/lan/wlan3_get_packet_statistics.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3GetBSSID](./pkg/tr064/lan/wlan3_get_b_ssid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3GetSSID](./pkg/tr064/lan/wlan3_get_ssid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3SetSSID](./pkg/tr064/lan/wlan3_set_ssid.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3GetBeaconType](./pkg/tr064/lan/wlan3_get_beacon_type.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3SetBeaconType](./pkg/tr064/lan/wlan3_set_beacon_type.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3GetChannelInfo](./pkg/tr064/lan/wlan3_get_channel_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3SetChannel](./pkg/tr064/lan/wlan3_set_channel.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3GetBeaconAdvertisement](./pkg/tr064/lan/wlan3_get_beacon_advertisement.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3SetBeaconAdvertisement](./pkg/tr064/lan/wlan3_set_beacon_advertisement.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3GetTotalAssociations](./pkg/tr064/lan/wlan3_get_total_associations.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3GetGenericAssociatedDeviceInfo](./pkg/tr064/lan/wlan3_get_generic_associated_device_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3GetSpecificAssociatedDeviceInfo](./pkg/tr064/lan/wlan3_get_specific_associated_device_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3XavmGetSpecificAssociatedDeviceInfoByIp](./pkg/tr064/lan/wlan3_xavm_get_specific_associated_device_info_by_ip.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3XavmGetWLANDeviceListPath](./pkg/tr064/lan/wlan3_xavm_get_wlan_device_list_path.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3XavmSetStickSurfEnable](./pkg/tr064/lan/wlan3_xavm_set_stick_surf_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3XavmGetIPTVOptimized](./pkg/tr064/lan/wlan3_xavm_get_ip_tv_optimized.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3XavmSetIPTVOptimized](./pkg/tr064/lan/wlan3_xavm_set_ip_tv_optimized.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3XavmGetNightControl](./pkg/tr064/lan/wlan3_xavm_get_night_control.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3XavmGetWLANHybridMode](./pkg/tr064/lan/wlan3_xavm_get_wlan_hybrid_mode.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3XavmSetWLANHybridMode](./pkg/tr064/lan/wlan3_xavm_set_wlan_hybrid_mode.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3XavmGetWLANExtInfo](./pkg/tr064/lan/wlan3_xavm_get_wlan_ext_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3XavmGetWPSInfo](./pkg/tr064/lan/wlan3_xavm_get_wps_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3XavmSetWPSConfig](./pkg/tr064/lan/wlan3_xavm_set_wps_config.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3XavmSetWPSEnable](./pkg/tr064/lan/wlan3_xavm_set_wps_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3XavmSetWLANGlobalEnable](./pkg/tr064/lan/wlan3_xavm_set_wlan_global_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.Wlan3XavmGetWLANConnectionInfo](./pkg/tr064/lan/wlan3_xavm_get_wlan_connection_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.GetHostNumberOfEntries](./pkg/tr064/lan/get_host_number_of_entries.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.GetSpecificHostEntry](./pkg/tr064/lan/get_specific_host_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.GetGenericHostEntry](./pkg/tr064/lan/get_generic_host_entry.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmGetLanDeviceHostsInfo](./pkg/tr064/lan/xavm_get_lan_device_hosts_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmGetChangeCounter](./pkg/tr064/lan/xavm_get_change_counter.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmSetHostNameByMACAddress](./pkg/tr064/lan/xavm_set_host_name_by_mac_address.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmGetAutoWakeOnLANByMACAddress](./pkg/tr064/lan/xavm_get_auto_wake_on_lan_by_mac_address.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmSetAutoWakeOnLANByMACAddress](./pkg/tr064/lan/xavm_set_auto_wake_on_lan_by_mac_address.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmWakeOnLANByMACAddress](./pkg/tr064/lan/xavm_wake_on_lan_by_mac_address.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmGetSpecificHostEntryByIP](./pkg/tr064/lan/xavm_get_specific_host_entry_by_ip.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmHostsCheckUpdate](./pkg/tr064/lan/xavm_hosts_check_update.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmHostDoUpdate](./pkg/tr064/lan/xavm_host_do_update.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmSetPrioritizationByIP](./pkg/tr064/lan/xavm_set_prioritization_by_ip.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmGetHostListPath](./pkg/tr064/lan/xavm_get_host_list_path.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmGetMeshListPath](./pkg/tr064/lan/xavm_get_mesh_list_path.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmGetFriendlyName](./pkg/tr064/lan/xavm_get_friendly_name.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmSetFriendlyName](./pkg/tr064/lan/xavm_set_friendly_name.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmSetFriendlyNameByIP](./pkg/tr064/lan/xavm_set_friendly_name_by_ip.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.XavmSetFriendlyNameByMAC](./pkg/tr064/lan/xavm_set_friendly_name_by_mac.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.SetLanEthernetIfCfgEnable](./pkg/tr064/lan/set_lan_ethernet_if_cfg_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.GetLanEthernetIfCfgInfo](./pkg/tr064/lan/get_lan_ethernet_if_cfg_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.GetLanEthernetIfCfgStatistics](./pkg/tr064/lan/get_lan_ethernet_if_cfg_statistics.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.GetLanHCfgMgmInfo](./pkg/tr064/lan/get_lan_h_cfg_mgm_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.SetDHCPServerEnable](./pkg/tr064/lan/set_dhcp_server_enable.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.SetIPInterface](./pkg/tr064/lan/set_ip_interface.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.GetAddressRange](./pkg/tr064/lan/get_address_range.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.SetAddressRange](./pkg/tr064/lan/set_address_range.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.GetIPRoutersList](./pkg/tr064/lan/get_ip_routers_list.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.SetIPRouter](./pkg/tr064/lan/set_ip_router.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.GetSubnetMask](./pkg/tr064/lan/get_subnet_mask.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.SetSubnetMask](./pkg/tr064/lan/set_subnet_mask.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.GetDNSServers](./pkg/tr064/lan/get_dns_servers.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [lan.GetIPInterfaceNumberOfEntries](./pkg/tr064/lan/get_ip_interface_number_of_entries.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [wan.GetCommonLinkProperties](./pkg/tr064/wan/get_common_link_properties.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [wan.GetTotalBytesSent](./pkg/tr064/wan/get_total_bytes_sent.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [wan.GetTotalBytesReceived](./pkg/tr064/wan/get_total_bytes_received.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [wan.GetTotalPacketsSent](./pkg/tr064/wan/get_total_packets_sent.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [wan.GetTotalPacketsReceived](./pkg/tr064/wan/get_total_packets_received.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [wan.XavmSetWANAccessType](./pkg/tr064/wan/xavm_set_wan_access_type.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [wan.XavmGetActiveProvider](./pkg/tr064/wan/xavm_get_active_provider.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [wan.XavmGetOnlineMonitor](./pkg/tr064/wan/xavm_get_online_monitor.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [wan.GetWanDslIfConfigInfo](./pkg/tr064/wan/get_wan_dsl_if_config_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [wan.GetStatisticsTotal](./pkg/tr064/wan/get_statistics_total.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [wan.XavmGetDSLDiagnoseInfo](./pkg/tr064/wan/xavm_get_dsl_diagnose_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
| [wan.XavmGetDSLInfo](./pkg/tr064/wan/xavm_get_dsl_info.go) | FRITZ!Box 6490 Cable v141.07.57 |
<!-- MARKDOWN-AUTO-DOCS:END -->

## Additional library functions

The following functions offer comfort features, over the existing API functions.

| Go function                                                   | description                                            |
|---------------------------------------------------------------|--------------------------------------------------------|
| [lan.XavmGetHostList](./pkg/tr064/lan/hosts_get_host_list.go) | lists all known hosts, wraps `lan.XavmGetHostListPath` |
