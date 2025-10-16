# Event sourced design

<style type="text/css">
			.dynamotableviz-table {
				border-collapse: collapse;
				margin: 25px 0;
				font-size: 0.9em;
				font-family: sans-serif;
				min-width: 400px;
				border: solid 1px #dddddd;
				color: #000000;
			}
			th.dynamotableviz-pk {
				background-color: #D6EAF8;
				text-align: left;
				padding: 10px;
				font-weight: bolder;
				border: solid 1px #dddddd;
				color: #000000;
			}
			td.dynamotableviz-pk {
				text-align: left;
				padding: 10px;
				font-weight: bolder;
				border: solid 1px #dddddd;
				color: #000000;
			}
			th.dynamotableviz-sk {
				background-color: #EBF5FB;
				text-align: left;
				padding: 10px;
				font-weight: bolder;
				border: solid 1px #dddddd;
				color: #000000;
			}
			td.dynamotableviz-sk {
				text-align: left;
				padding: 10px;
				font-weight: bolder;
				border: solid 1px #dddddd;
				color: #000000;
			}
			th.dynamotableviz-key {
				background-color: #E9F7EF;
				text-align: left;
				padding: 10px;
				font-weight: bolder;
				border: solid 1px #dddddd;
				color: #000000;
			}
			td.dynamotableviz-key {
				text-align: left;
				padding: 10px;
				border: solid 1px #dddddd;
				color: #000000;
			}
			th.dynamotableviz-attr {
				background-color: #eeeeee;
				text-align: left;
				padding: 10px;
				font-weight: bolder;
				border: solid 1px #dddddd;
				color: #000000;
			}
			td.dynamotableviz-attr {
				text-align: left;
				padding: 10px;
				border: solid 1px #dddddd;
				color: #000000;
			}
			tr.dynamotableviz-even {
				background-color: #ffffff;
				color: #000000;
			}
			tr.dynamotableviz-odd {
				background-color: #eeeeee;
				color: #000000;
			}
		</style>

<Transform :scale="0.8">
	<table class="dynamotableviz-table"><tbody><tr><th class="dynamotableviz-pk">pk</th><th class="dynamotableviz-sk">sk</th><th class="dynamotableviz-key">_seq</th><th class="dynamotableviz-key">_type</th><th class="dynamotableviz-key">_date</th><th class="dynamotableviz-attr" colspan="3">Attributes</th></tr><tr class="dynamotableviz-even"><td class="dynamotableviz-pk" rowspan="5">account/1234</td><td class="dynamotableviz-sk">INBOUND/1/0/Transfer</td><td class="dynamotableviz-key">1</td><td class="dynamotableviz-key">PayIn</td><td class="dynamotableviz-key">2023-01-01</td><td class="dynamotableviz-attr">txType=&#34;IBAN&#34;</td><td class="dynamotableviz-attr">amt=&#34;5000&#34;</td><td class="dynamotableviz-attr">desc=&#34;From Dave&#34;</td></tr><tr class="dynamotableviz-even"><td class="dynamotableviz-sk">INBOUND/2/0/Transaction</td><td class="dynamotableviz-key">2</td><td class="dynamotableviz-key">CardPayment</td><td class="dynamotableviz-key">2023-01-02</td><td class="dynamotableviz-attr">txType=&#34;Visa&#34;</td><td class="dynamotableviz-attr">amt=&#34;1000&#34;</td><td class="dynamotableviz-attr">retailer=&#34;Tesco&#34;</td></tr><tr class="dynamotableviz-even"><td class="dynamotableviz-sk">INBOUND/3/0/Transaction</td><td class="dynamotableviz-key">3</td><td class="dynamotableviz-key">CardPayment</td><td class="dynamotableviz-key">2023-01-03</td><td class="dynamotableviz-attr">txType=&#34;Visa&#34;</td><td class="dynamotableviz-attr">amt=&#34;6000&#34;</td><td class="dynamotableviz-attr">retailer=&#34;Brewdog&#34;</td></tr><tr class="dynamotableviz-even"><td class="dynamotableviz-sk">OUTBOUND/3/0</td><td class="dynamotableviz-key">3</td><td class="dynamotableviz-key">Notification</td><td class="dynamotableviz-key">2023-01-03</td><td class="dynamotableviz-attr">msg=&#34;You have gone overdrawn&#34;</td><td class="dynamotableviz-attr" colspan="2"></td></tr><tr class="dynamotableviz-even"><td class="dynamotableviz-sk">STATE</td><td class="dynamotableviz-key">3</td><td class="dynamotableviz-key">Account</td><td class="dynamotableviz-key">2023-01-03</td><td class="dynamotableviz-attr">balance=&#34;-1000&#34;</td><td class="dynamotableviz-attr">name=&#34;Current account&#34;</td><td class="dynamotableviz-attr" colspan="1"></td></tr></tbody></table>
		</Transform>
