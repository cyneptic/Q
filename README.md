This project is a SuperApp containing multiple applications and microservices within those applications;

1. letsgo_authentication; authentication service for the entire superapp, including user creation, login and logout services, admin creation and role assignment along with other authentication services used within different applications

2. letsgo_ticket_mock; this is a service mocking an actual flight provider, it generates and mocks flight information for over 3 years of mocked data.

3. letsgo_ticketpanel; this is the entire back-end structure to an application designed to sell tickets to flights provided by different flight providers, its integrated with a postgres database and a fake payment gateway by banktest.ir

4. letsgo_sms_mock; this is a service mocking different devices receiving messages from our sms panel, it receives information of a message including source, recipient and the content of messages and stores it in a postgres database.

5. letsgo_smspanel; this is the entire back-end structure to a sms panel application, users can buy or subscribe to numbers, make contacts and contactlists, send messages to their contacts via number or nicknames, and even send messages on an interval to their contactlists, it includes an admin panel back-end granting the smspanel admins access to functions such as changing the prices and searching through sent messages, it implements a postgres database, and kavenegar's sms service.

- this application was designed by lets_go team as a final_project to Quera Back-End Engineering with Golang bootcamp -
