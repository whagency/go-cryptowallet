package cryptowallet

var (
	TestResponseGetCurrencies   = "{\"status\":\"OK\",\"data\":[{\"currencyCode\":\"TEST\",\"tokenId\":null,\"fractionNumber\":8,\"enabled\":false},{\"currencyCode\":\"BPS\",\"tokenId\":null,\"fractionNumber\":8,\"enabled\":false},{\"currencyCode\":\"USDT\",\"tokenId\":null,\"fractionNumber\":8,\"enabled\":true},{\"currencyCode\":\"XRP\",\"tokenId\":null,\"fractionNumber\":6,\"enabled\":true},{\"currencyCode\":\"WAVES\",\"tokenId\":null,\"fractionNumber\":8,\"enabled\":false},{\"currencyCode\":\"EOS\",\"tokenId\":null,\"fractionNumber\":4,\"enabled\":false},{\"currencyCode\":\"XRP_TOKEN\",\"tokenId\":null,\"fractionNumber\":6,\"enabled\":false},{\"currencyCode\":\"ETH\",\"tokenId\":null,\"fractionNumber\":18,\"enabled\":true},{\"currencyCode\":\"ETC\",\"tokenId\":null,\"fractionNumber\":18,\"enabled\":false},{\"currencyCode\":\"DASH\",\"tokenId\":null,\"fractionNumber\":8,\"enabled\":false},{\"currencyCode\":\"BCH\",\"tokenId\":null,\"fractionNumber\":8,\"enabled\":false},{\"currencyCode\":\"QTUM\",\"tokenId\":null,\"fractionNumber\":8,\"enabled\":false},{\"currencyCode\":\"NEO\",\"tokenId\":null,\"fractionNumber\":0,\"enabled\":false},{\"currencyCode\":\"IEX\",\"tokenId\":null,\"fractionNumber\":4,\"enabled\":false},{\"currencyCode\":\"BTC\",\"tokenId\":null,\"fractionNumber\":8,\"enabled\":true},{\"currencyCode\":\"LTC\",\"tokenId\":null,\"fractionNumber\":8,\"enabled\":true},{\"currencyCode\":\"BNB\",\"tokenId\":null,\"fractionNumber\":18,\"enabled\":true},{\"currencyCode\":\"BNB_TOKEN\",\"tokenId\":null,\"fractionNumber\":0,\"enabled\":true},{\"currencyCode\":\"ADA\",\"tokenId\":null,\"fractionNumber\":6,\"enabled\":false},{\"currencyCode\":\"OMNI\",\"tokenId\":null,\"fractionNumber\":8,\"enabled\":false},{\"currencyCode\":\"TRX_TOKEN\",\"tokenId\":null,\"fractionNumber\":0,\"enabled\":true},{\"currencyCode\":\"XMR\",\"tokenId\":null,\"fractionNumber\":12,\"enabled\":false},{\"currencyCode\":\"TRX\",\"tokenId\":null,\"fractionNumber\":6,\"enabled\":true},{\"currencyCode\":\"BTG\",\"tokenId\":null,\"fractionNumber\":8,\"enabled\":false},{\"currencyCode\":\"ETH_TOKEN\",\"tokenId\":null,\"fractionNumber\":0,\"enabled\":true},{\"currencyCode\":\"ZEK\",\"tokenId\":null,\"fractionNumber\":8,\"enabled\":false},{\"currencyCode\":\"XLM\",\"tokenId\":null,\"fractionNumber\":8,\"enabled\":false}]}"
	TestResponseGetTokens       = "{\"status\":\"OK\",\"data\":[{\"currencyCode\":\"ALTC\",\"tokenId\":\"0x9Ab8c9CAcfD80CE623102EB68fD714f57cA5C182\",\"fractionNumber\":4,\"enabled\":true},{\"currencyCode\":\"ALFX\",\"tokenId\":\"0x5bfc839f5b8c1567e781dec5247db07734624659\",\"fractionNumber\":4,\"enabled\":true},{\"currencyCode\":\"USDT\",\"tokenId\":\"0xdac17f958d2ee523a2206206994597c13d831ec7\",\"fractionNumber\":6,\"enabled\":true},{\"currencyCode\":\"3-Ex\",\"tokenId\":\"0xa916D786e90d81DC7E41751d0201C87548A231D4\",\"fractionNumber\":0,\"enabled\":true}]}"
	TestResponseAddAddress      = "{\"status\":\"OK\",\"data\":{\"hash\":\"test5UvvGgf5942yfdPxM4mWiuQW4wkCQG4UG\"}}"
	TestResponseGetBalance      = "{\"status\":\"OK\",\"data\":{\"walletId\":\"test5UvvGgf5942yfdPxM4mWiuQW4wkCQG4UG\",\"balance\":1200000}}"
	TestResponseGetTransactions = "{\"status\":\"OK\",\"data\":{\"content\":[{\"id\":2389,\"timestamp\":\"2022-11-20T18:58:47Z\",\"timestampLong\":1668970727000,\"currency\":\"ETH_TOKEN::USDT::0xdac17f958d2ee523a2206206994597c13d831ec7\",\"amount\":110000000,\"type\":\"IN\",\"sourceAddresses\":[\"0x53c65018a7b15334f268d57c593306696018a404\"],\"destinationAddress\":\"0x2fdfcbf9aee86f53f7edb096d0f34696e9029d0a\",\"ipAddress\":\"\"},{\"id\":2369,\"timestamp\":\"2022-10-31T11:22:35Z\",\"timestampLong\":1667215355000,\"currency\":\"ETH\",\"amount\":50000000000000000,\"type\":\"IN\",\"sourceAddresses\":[\"0x92f9606fdc8c2b7cd1a3b341b65368bac0f5a80b\"],\"destinationAddress\":\"0x420fe57eed05fdfe9c337d145f42af897942873a\",\"ipAddress\":\"\"},{\"id\":2363,\"timestamp\":\"2022-10-25T11:51:23Z\",\"timestampLong\":1666698683000,\"currency\":\"ETH_TOKEN::USDT::0xdac17f958d2ee523a2206206994597c13d831ec7\",\"amount\":10000000,\"type\":\"IN\",\"sourceAddresses\":[\"0x92f9606fdc8c2b7cd1a3b341b65368bac0f5a80b\"],\"destinationAddress\":\"0x3674e07f21c26fef43503df4b82efa52b9565e2a\",\"ipAddress\":\"\"},{\"id\":2362,\"timestamp\":\"2022-10-24T09:00:59Z\",\"timestampLong\":1666602059000,\"currency\":\"ETH_TOKEN::USDT::0xdac17f958d2ee523a2206206994597c13d831ec7\",\"amount\":10000000,\"type\":\"IN\",\"sourceAddresses\":[\"0x92f9606fdc8c2b7cd1a3b341b65368bac0f5a80b\"],\"destinationAddress\":\"0x9db3a2b4d83566584e12ca5badd8ef9b94547d6c\",\"ipAddress\":\"\"},{\"id\":2200,\"timestamp\":\"2022-08-03T14:49:02Z\",\"timestampLong\":1659538142000,\"currency\":\"BTC\",\"amount\":47165,\"type\":\"IN\",\"sourceAddresses\":[\"bc1qaqjcqnlxq05sh53rm7flwcennvsphq8qpmgx4w\"],\"destinationAddress\":\"3ASubnmR9uKGjSWe4AGBhPghHkhtAatLTn\",\"ipAddress\":\"\"},{\"id\":2182,\"timestamp\":\"2022-07-15T14:00:47Z\",\"timestampLong\":1657893647000,\"currency\":\"BTC\",\"amount\":10000,\"type\":\"IN\",\"sourceAddresses\":[\"bc1qjwk2h2urxgk3epeh9k6nz6x5kp6h3yn4g0frcw\"],\"destinationAddress\":\"31yhThEhVwqvxwgm3Hv5CjUddtbZB2y7EN\",\"ipAddress\":\"\"},{\"id\":2133,\"timestamp\":\"2022-06-14T07:39:05Z\",\"timestampLong\":1655192345000,\"currency\":\"BNB_TOKEN::USDC::0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d\",\"amount\":9710000000000000000,\"type\":\"IN\",\"sourceAddresses\":[\"0xe2fc31f816a9b94326492132018c3aecc4a93ae1\"],\"destinationAddress\":\"0x536429dc6330ca6e0c3a2313dd6f5887dfb3605a\",\"ipAddress\":\"\"},{\"id\":2127,\"timestamp\":\"2022-06-14T03:15:41.091Z\",\"timestampLong\":1655176541091,\"currency\":\"TRX_TOKEN::USDT::TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t\",\"amount\":1500000000,\"type\":\"IN\",\"sourceAddresses\":[\"TQ7wK19fhZZqLdj2Xcw2e6Ejs3cTZbfBbF\"],\"destinationAddress\":\"TK7gjDNSm2iP2f69acGEVWVxTjEzvMQAfY\",\"ipAddress\":\"\"},{\"id\":2111,\"timestamp\":\"2022-06-10T11:49:33Z\",\"timestampLong\":1654861773000,\"currency\":\"BNB_TOKEN::USDC::0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d\",\"amount\":9710000000000000000,\"type\":\"IN\",\"sourceAddresses\":[\"0x8894e0a0c962cb723c1976a4421c95949be2d4e3\"],\"destinationAddress\":\"0xfbfcd8773f136553e64b2495a01ceaefc12f7502\",\"ipAddress\":\"\"}],\"totalPages\":1,\"totalElements\":9,\"size\":100,\"page\":0}}"
)
