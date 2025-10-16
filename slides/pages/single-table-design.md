# Single table design

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

<Transform :scale="0.8" origin="top center">
<table class="dynamotableviz-table"><tbody><tr><th class="dynamotableviz-pk">pk</th><th class="dynamotableviz-sk">sk</th><th class="dynamotableviz-key">purchaseDate</th><th class="dynamotableviz-key">expires</th><th class="dynamotableviz-key">type</th><th class="dynamotableviz-attr" colspan="1">Attributes</th></tr><tr class="dynamotableviz-even"><td class="dynamotableviz-pk" rowspan="3">email/test@example.com</td><td class="dynamotableviz-sk">policy/home/5678</td><td class="dynamotableviz-key">2023-02-01</td><td class="dynamotableviz-key">2024-02-01</td><td class="dynamotableviz-key">home</td><td class="dynamotableviz-attr" colspan="1"></td></tr><tr class="dynamotableviz-even"><td class="dynamotableviz-sk">policy/home/9012</td><td class="dynamotableviz-key">2023-02-02</td><td class="dynamotableviz-key">2024-02-02</td><td class="dynamotableviz-key">home</td><td class="dynamotableviz-attr" colspan="1"></td></tr><tr class="dynamotableviz-even"><td class="dynamotableviz-sk">policy/vehicle/1234</td><td class="dynamotableviz-key">2023-01-01</td><td class="dynamotableviz-key">2024-01-01</td><td class="dynamotableviz-key">vehicle</td><td class="dynamotableviz-attr">reg=&#34;KW12DFF&#34;</td></tr><tr class="dynamotableviz-odd"><td class="dynamotableviz-pk" rowspan="2">email/other@example.com</td><td class="dynamotableviz-sk">policy/vehicle/5678</td><td class="dynamotableviz-key">2023-01-01</td><td class="dynamotableviz-key">2024-01-01</td><td class="dynamotableviz-key">vehicle</td><td class="dynamotableviz-attr">reg=&#34;VS12HGD&#34;</td></tr><tr class="dynamotableviz-odd"><td class="dynamotableviz-sk">policy/vehicle/5678</td><td class="dynamotableviz-key">2023-03-01</td><td class="dynamotableviz-key">2024-03-01</td><td class="dynamotableviz-key">vehicle</td><td class="dynamotableviz-attr">reg=&#34;B614XRF&#34;</td></tr></tbody></table>
</Transform>
