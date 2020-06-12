package main

import (
	"fmt"
	"net/http"
	"runtime"
	"flag"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<!DOCTYPE html>
<!-- saved from url=(0017)https://bmstu.ru/ -->
<html lang="ru"><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<meta name="google-site-verification" content="ArHTqIMCyLyjpyavw-el7KLwmmXJ3_x2ZbJzEfAQDOw">
<title>МГТУ им. Н.&nbsp;Э.&nbsp;Баумана</title>
<meta http-equiv="Content-Security-Policy" content="upgrade-insecure-requests">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta http-equiv="expires" content="Tue, 01 Jan 1980 1:00:00 GMT">

<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="description" content="Официальный сайт Московского Государственного Технического Университета имени Н. Э. Баумана">
<meta http-equiv="expires" content="Tue, 01 Jan 1980 1:00:00 GMT">
<!-- Latest compiled and minified CSS -->
<link rel="stylesheet" href="./МГТУ им. Н. Э. Баумана_files/bootstrap.min.css">
<link rel="stylesheet" media="all" href="./МГТУ им. Н. Э. Баумана_files/layout.css">
<link rel="stylesheet" media="all" href="./МГТУ им. Н. Э. Баумана_files/shared.css">
<link rel="stylesheet" media="all" href="./МГТУ им. Н. Э. Баумана_files/bs-customization.css">
<!-- Begin Temp for Priemnaya company -->
<!-- <link rel="stylesheet" type="text/css" href="//cdn.jsdelivr.net/jquery.slick/1.6.0/slick.css"/> -->
<link rel="stylesheet" type="text/css" href="./МГТУ им. Н. Э. Баумана_files/slick.css">
<!-- End Temp for Priemnaya company -->
<link rel="stylesheet" media="all" href="./МГТУ им. Н. Э. Баумана_files/slick-theme.css">

<!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
<!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
<!--[if lt IE 9]>
<link rel="stylesheet" media="all" href="/b-twig/assets/css/ie8-.css" >
<script type="text/javascript" src="/b-twig/assets/libs/html5shiv.js"></script>
<script type="text/javascript" src="/b-twig/assets/libs/respond.min.js"></script>
<![endif]-->
<!--[if gt IE 8]>
<link rel="stylesheet" media="all" href="/b-twig/assets/css/ie9+.css" >
<!--<![endif]-->
<link rel="shortcut icon" href="https://bmstu.ru/favicon/favicon-2.ico" type="image/x-icon">
<meta name="sputnik-verification" content="157UWyLH6Qhcj1Ei">
<style id="holderjs-style" type="text/css"></style></head>
<body>
<div id="message" class="j-hidden"></div>
<div id="mainbody">
	
	
	<div id="navbar">
		<div class="topbar topbar-desktop bg-shadow">
			<div class="container">
				<div class="row">
					<div class="col-md-12">
						<div class="homepage-emblem-offset">
							<a class="emblem desktop" href="https://bmstu.ru/"><img src="./МГТУ им. Н. Э. Баумана_files/bmstu-emblem.png" alt=""></a>
							<div class="topbar-body offset">
								<div class="topbar-body">
									<div class="emblem-bg"></div>
									<div class="topbar-decor"></div>									
									<ul class="nav-extra pull-right">
										<li><a class="j-color-red" href="http://www.bmstu.ru/en/">English</a></li>
										<li><a class="hidden-xs" href="https://bmstu.ru/mstu/contacts/map/">Контакты</a></li>
										<li><a href="http://www.bmstu.ru/plain/" class="hidden-xs" itemprop="Copy">Для слабовидящих</a></li>
										<li><a class="hidden-xs" href="http://www.bmstu.ru/feedback/">Обращения граждан</a></li>
										<li><a class="hidden-xs" href="http://www.bmstu.ru/mstu/ya-search/">Поиск</a></li>
									</ul>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
		<div class="j-hidden">
			<div class="container">
				<div class="row">
					<div class="col-xs-2">
						<a class="emblem mobile" href="https://bmstu.ru/"><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/bmstu-emblem.png" alt=""></a>
					</div>
					<div class="col-xs-10">			
						<a class="pull-right j-display-iblock j-padd-top-sm j-padd-left j-padd-bottom j-color-red j-larger j-bold" href="http://www.bmstu.ru/en/">English</a>
					</div>
				</div>
			</div>			
		</div>
		<div id="navbar-body">
			<div class="container">
				<div class="row">
					<div class="col-xs-12 col-md-12">
						<div class="homepage-emblem-offset">
							<div class="navbar-offset">

								<nav class="navbar navbar-bauman" role="navigation">
									<!-- Brand and toggle get grouped for better mobile display -->
									<div class="navbar-header">
									<button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
										<span class="sr-only">Toggle navigation</span>Меню
									</button>
									<a class="visible-xs navbar-brand" href="https://bmstu.ru/">
										<span id="navbar-brand-part-1">МГТУ</span>
										<span id="navbar-brand-part-2"><span id="navbar-brand-part-3">имени&nbsp;Н.Э</span>Баумана</span>
									</a>
									</div>

									<!-- Collect the nav links, forms, and other content for toggling -->
									<div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
										<ul class="nav navbar-nav nav-justified">
										<li class="dropdown"><a href="https://bmstu.ru/#" class="dropdown-toggle" data-toggle="dropdown"><span class="j-navbar-tip">Поступающим</span> <b class="caret"></b></a><ul class="dropdown-menu"><li><a href="http://www.bmstu.ru/abitur">Абитуриент</a></li><li><a href="http://www.bmstu.ru/master">Магистрант</a></li><li><a href="https://abiturient.bmstu.ru/">День открытых дверей</a></li><li><a href="http://www.bmstu.ru/mstu/works/edu/#space-college">Техникум 9-11 класс</a></li><li><a href="https://bmstu.ru/mstu/admissions/postgraduate/">Аспирантура и Докторантура</a></li><li><a href="http://isot.bmstu.ru/2edu.html">Второе высшее</a></li><li><a href="http://edu.bmstu.ru/" target="_blank">Образовательный центр</a></li><li><a href="http://isot.bmstu.ru/pk">Повышение квалификации</a></li><li><a href="http://military.bmstu.ru/">Военный учебный центр</a></li><li><a href="https://bmstu.ru/mstu/admissions/preparation/">Подготовка</a></li><li><a href="https://bmstu.ru/mstu/admissions/transfer/">Восстановление и переводы</a></li><li><a href="https://inginirium.ru/" target="_blank">Курсы для школьников</a></li></ul></li><li class="dropdown"><a href="https://bmstu.ru/#" class="dropdown-toggle" data-toggle="dropdown">Обучающимся <b class="caret"></b></a><ul class="dropdown-menu"><li><a href="https://bmstu.ru/mstu/undergraduate/exam-schedule/">Расписание экзаменов</a></li><li><a href="https://bmstu.ru/mstu/undergraduate/schedule/">Расписание занятий</a></li><li><a href="https://bmstu.ru/mstu/undergraduate/studies/">Общая информация</a></li><li><a href="https://bmstu.ru/mstu/undergraduate/juniors/">1, 2 курсы</a></li><li><a href="https://bmstu.ru/mstu/undergraduate/seniors/">Старшие курсы</a></li><li><a href="https://bmstu.ru/mstu/undergraduate/post-graduates/">Аспирант</a></li><li><a href="http://www.mtkp.ru/student/">Техникум</a></li><li><a href="https://bmstu.ru/mstu/undergraduate/returning/">Восстановление</a></li><li><a href="http://library.bmstu.ru/">Библиотека</a></li><li><a href="https://bmstu.ru/mstu/undergraduate/sopk/">Работа в студенческом отряде «Приёмная комиссия»</a></li></ul></li><li class="dropdown"><a href="https://bmstu.ru/#" class="dropdown-toggle" data-toggle="dropdown"><span class="hidden-xs hidden-sm navbaritem-decor"></span>Университет <b class="caret"></b></a><ul class="dropdown-menu"><li><a href="https://bmstu.ru/sveden">Сведения об образовательной организации</a></li><li><a href="https://bmstu.ru/mstu/about/head/">Ректор</a></li><li><a href="https://bmstu.ru/mstu/about/organization/">Структура Университета</a></li><li><a href="https://bmstu.ru/mstu/about/introducing/">История</a></li><li><a href="https://bmstu.ru/mstu/about/trustees-board/">Попечительский совет</a></li><li class="divider"></li><li><a href="https://bmstu.ru/mstu/about/senate/">Ученый совет</a></li><li><a href="http://www.bmstu.ru/mstu/about/staff/">Сотрудникам</a></li><li><a href="https://bmstu.ru/mstu/about/graduate/">Выпускникам</a></li><li><a href="https://bmstu.ru/mstu/about/documents/">Нормативные документы и приказы</a></li><li><a href="https://bmstu.ru/mstu/about/anticorr/">Противодействие коррупции</a></li></ul></li><li class="dropdown"><a href="https://bmstu.ru/#" class="dropdown-toggle" data-toggle="dropdown">Деятельность <b class="caret"></b></a><ul class="dropdown-menu"><li><a href="https://bmstu.ru/mstu/works/edu/">Учебная</a></li><li class="divider"></li><li><a class="j-textblock" href="https://bmstu.ru/mstu/works/science/">Научная<br><span class="j-smaller-x">Национальный Исследовательский Университет</span></a></li><li class="divider"></li><li><a href="https://bmstu.ru/mstu/works/enterprise/">Коммерческая</a></li><li><a href="https://bmstu.ru/mstu/works/international/">Международная</a></li><li><a href="https://bmstu.ru/mstu/works/pr-media/">Информационная</a></li><li><a href="https://bmstu.ru/mstu/works/publishing-activities/">Издательская</a></li><li><a href="https://bmstu.ru/mstu/works/social/">Общественная</a></li></ul></li><li class="dropdown"><a href="https://bmstu.ru/#" class="dropdown-toggle" data-toggle="dropdown">Активность <b class="caret"></b></a><ul class="dropdown-menu"><li><a href="http://culture.bmstu.ru/">Дворец культуры</a></li><li><a href="https://bmstu.ru/mstu/activity/university-exhibitions/">Выставки</a></li><li><a href="http://sport.bmstu.ru/">Спорт</a></li><li><a href="http://www.bmstu.ru/ps/">Персональные страницы</a></li><li><a href="http://www.bmstu.ru/ps/forum">Форум</a></li><li><a href="http://mail.bmstu.ru/">Электронная почта</a></li><li><a href="https://bmstu.ru/mstu/activity/support/">Редакция сайта</a></li><li><a href="https://bmstu.ru/mstu/antiterror/">Антитеррор</a></li></ul></li>											
										</ul>
									</div><!-- /.navbar-collapse -->
								</nav>								
								
								
							</div>
						</div>
					</div>
				</div>		
			</div>
		</div>		
	</div>
	
	
	<div id="main-carousel" class="hidden-xs carousel slide carousel-bauman" data-ride="carousel">
		<!-- Штатные Indicators -->
		
		<ol class="carousel-indicators">
			<li data-target="#main-carousel" data-slide-to="0" class="active"></li>
			<li data-target="#main-carousel" data-slide-to="1" class=""></li>
			<li data-target="#main-carousel" data-slide-to="2" class=""></li>
		</ol>
		
		<!-- End. Штатные Wrapper for slides -->
		<div class="carousel-inner">

			<div class="item active" id="carousel-start-item" style="background: none;">
				<div class="container" style="width: 1320px;">
					<div class="row hidden-xs">

						<div class="item"><center><a href="https://bmstu.ru/mstu/accreditation/"><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/banner-accreditation.jpg" alt="Информация для экспертов по аккредитации МГТУ им. Н.Э. Баумана в 2020 г."></a></center></div>




					</div>

				</div>

			</div>
			

			
			<div class="item"><center><a href="https://bmstu.ru/mstu/undergraduate/coronavirus/"><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/banner_0531.png" alt="Дистанционное обучение"></a></center></div>
			<div class="item"><center><a href="https://bmstu.ru/mstu/undergraduate/sopk/" target="_blank"><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/banner-sopk2.png" alt="Работа в студенческом отряде «Приёмная комиссия»"></a></center></div>		
		</div>

		<!-- Controls -->
		<!-- Штатные Controls -->
		<a class="left carousel-control" href="https://bmstu.ru/#main-carousel" data-slide="prev" style="width: 318px;">
			<center class="j-valign-center" style="padding-top: 120.5px;">
				<span class="carousel-nav-arrow"><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/carousel-navarrow-left.png"></span>
				<span class="glyphicon glyphicon-chevron-left"></span>
			</center>
		</a>
		<a class="right carousel-control" href="https://bmstu.ru/#main-carousel" data-slide="next" style="width: 318px;">
			<center class="j-valign-center" style="padding-top: 120.5px;">
				<span class="carousel-nav-arrow"><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/carousel-navarrow-right.png"></span>
				<span class="glyphicon glyphicon-chevron-right"></span>
			</center>
		</a>
		<!-- End. Штатные Controls -->
	</div>


	
	<div class="hidden-xs container">
		<div class="row j-padd2">
			<div class="col-md-12">

			<h3 class="text-center"><a href="https://bmstu.ru/mstu/undergraduate/coronavirus/"><strong><span style="color: #AA0000 !important;">Дистанционное обучение</span></strong></a></h3>
			</div>
		</div>	
	</div>
	



	<div class="hidden-xs container">
		<div class="row j-padd2">
			<div class="col-md-12">


			<center><a href="https://bmstu.ru/mstu/about/introducing/"><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/awards.png" alt=""></a></center>
		
			</div>
		</div>	
	</div>


	<!--
	<div id="9may" class="container">
		<div class="row">
			<div class="col-md-3"><div style="position:relative; height:102px;"><div style="position: absolute; bottom: 0;"><img src="http://www.bmstu.ru/content/shared/img/9may.gif"></div></div></div>
			<div class="col-md-9 "><h1 style="color:orange; font-style:italic; text-align:right; font-weight:bold; padding-top:10px;" class="">С Днем Победы! Я помню... я горжусь</h1></div>
			
		</div>	
	</div>
	-->



	<div class="hidden-xs j-border-top j-bg-white">
		<div class="container">
			<div class="row j-padd-bottom-lg j-padd-top-lg j-textblock">
				<div class="col-xs-3 col-sm-3 col-md-3">
					<!-- <a class="j-uppercase j-bold j-icon-placeholder" href="http://www.bmstu.ru/mstu/admissions/"><span class="j-icon j-icon-admission"></span>Поступающим</a> -->
					<p class="hidden-xs j-uppercase j-bold j-icon-placeholder"><span class="j-icon j-icon-admission"></span>Поступающим</p>
					<p class="hidden-xs j-padd-top-sm"><small class="j-color-gray">как поступить, справочники, правила приёма, схема поступления, подготовка, график и план приёма, льготы, проходной балл</small></p>
					<p class="hidden-xs"><a href="http://www.bmstu.ru/abitur/" class="btn btn-info btn-xs">Абитуриенту</a>&nbsp;<a href="http://www.bmstu.ru/master/" class="btn btn-info btn-xs">Магистранту</a></p>
				</div>
				<div class="col-xs-3 col-sm-3 col-md-3">
					<a class="j-uppercase j-bold j-icon-placeholder" href="https://bmstu.ru/mstu/undergraduate/studies"><span class="j-icon j-icon-student"></span>Студенту <!--<sup style="text-transform:lowercase;" class="j-color-yellow">new</sup>--></a> 
					<p class="j-padd-top-sm"><small class="j-color-gray">расписание, стипендии, студенческие клубы и сообщества, кабинет студента, профком студентов, путёвки, конкурсы</small></p>
					<p><a href="https://bmstu.ru/mstu/undergraduate/studies" class="btn btn-info btn-xs">Перейти</a></p>
				</div>
				<div class="col-xs-3 col-sm-3 col-md-3">
					<a class="j-uppercase j-bold j-icon-placeholder" href="https://bmstu.ru/mstu/about/graduate/"><span class="j-icon j-icon-alumni"></span>Выпускнику</a>
					<p class="j-padd-top-sm"><small class="j-color-gray">помощь в трудоустройстве, второе высшее образование, аспирантура, гранты, клуб выпускников</small></p>
					<p><a href="https://bmstu.ru/mstu/about/graduate/" class="btn btn-info btn-xs">Перейти</a></p>				
				</div>
				<div class="col-xs-3 col-sm-3 col-md-3">
					<a class="j-uppercase j-bold j-icon-placeholder" href="https://bmstu.ru/mstu/about/staff/"><span class="j-icon j-icon-staff"></span>Сотруднику</a>
					<p class="j-padd-top-sm"><small class="j-color-gray">персональные страницы преподавателей и сотрудников, доска почёта, повышение квалификации, профсоюз</small></p>
					<p><a href="https://bmstu.ru/mstu/about/staff/" class="btn btn-info btn-xs">Перейти</a></p>
				</div>

			</div>	
		</div>
	</div>
	<div class="visible-xs">

		<div class="container">
			<div class="row">
				<div class="list-group">
										<a href="https://bmstu.ru/mstu/undergraduate/coronavirus/" class="list-group-item">
						<h4 class="list-group-item-heading">ВАЖНО!</h4>
						<p class="list-group-item-text"><strong>Дистанционное обучение</strong></p>
					</a>
										<a href="http://www.bmstu.ru/abitur/" class="list-group-item">
						<h4 class="list-group-item-heading">Абитуриенту</h4>
						<p class="list-group-item-text">как поступить, справочники, правила приёма, схема поступления, подготовка, график и план приёма, льготы, проходной балл</p>
					</a>
					<a href="http://www.bmstu.ru/master/" class="list-group-item">
						<h4 class="list-group-item-heading">Магистранту</h4>
						<p class="list-group-item-text">как поступить, справочники, правила приёма, схема поступления, подготовка, график и план приёма, льготы, проходной балл</p>
					</a>
					<a href="https://bmstu.ru/mstu/undergraduate/studies" class="list-group-item">
						<h4 class="list-group-item-heading">Студенту</h4>
						<p class="list-group-item-text">расписание, стипендии, студенческие клубы и сообщества, кабинет студента, профком студентов, путёвки, конкурсы</p>
					</a>
					<a href="https://bmstu.ru/mstu/about/graduate/" class="list-group-item">
						<h4 class="list-group-item-heading">Выпускнику</h4>
						<p class="list-group-item-text">помощь в трудоустройстве, второе высшее образование, аспирантура, гранты, клуб выпускников</p>
					</a>
					<a href="https://bmstu.ru/mstu/about/staff/" class="list-group-item">
						<h4 class="list-group-item-heading">Сотруднику</h4>
						<p class="list-group-item-text">персональные страницы преподавателей и сотрудников, доска почёта, повышение квалификации, профсоюз</p>
					</a>		
					<a href="http://www.bmstu.ru/plain/" class="list-group-item" itemprop="Copy">
						<h4 class="list-group-item-heading">Для слабовидящих</h4>
						<p class="list-group-item-text">версия сайта для слабовидящих граждан</p>
					</a>				
				</div>
			</div>
		</div>

	</div>
	
	<div class="hidden-xs popular-links">
		<div class="popular-links-bg-right"></div>
		<div class="visible-md visible-lg popular-links-bg-left"></div>
		<div class="popular-links-body">
			<div class="container">

				<div class="row j-padd-top-sm j-padd-bottom-sm">
					
				
					<div class="col-sm-4 col-md-4">
						<table><tbody><tr>
						<td class="visible-lg "><h3 class="j-marg-sm j-bold" style="font-size: 18px;"><i>Полезное</i></h3></td>
						<td class="visible-lg"><img src="./МГТУ им. Н. Э. Баумана_files/decor-brace.png"></td>
						<td class="j-padd-left-sm">
							<ul class="j-vlist j-smaller j-padd-0">
							<li class="j-marg-bottom-xs"><a href="https://bmstu.ru/mstu/about/organization/">Факультеты, Кафедры, Филиалы</a></li>
							<!-- <li class="j-marg-bottom-xs"><a href="/mstu/about/university/divisions/">Подразделения и Управления</a></li> -->
							<li class="j-marg-bottom-xs"><a href="https://bmstu.ru/mstu/about/university/documents/">Нормативные документы, Бланки</a></li>
							<li class="j-marg-bottom-xs"><a href="https://bmstu.ru/abitur/general/qanda/">Часто задаваемые вопросы</a></li>
							<li class="j-marg-bottom-xs"><a href="https://bmstu.ru/mstu/about/organization#research-centers">Научные центры, ЦКП и УНУ</a></li>
							<!--<li class="j-marg-bottom-xs"><a href="http://usu-beam.bmstu.ru/" target="_blank">Уникальные научные установки и Центры Коллективного Пользования</a></li>-->
										
							<!-- <li class="j-marg-bottom-xs"><a href="/mstu/about/introducing/bauman-glossary/">Бауманский глоссарий</a> </li> -->
							<!-- <li class="j-marg-bottom-xs"><a href="/mstu/about/introducing/virtual-tour/">Виртуальная прогулка</a></li> -->
							<li class="j-marg-bottom-xs"><a href="http://shop.bmstu.ru/" target="_blank">Магазин атрибутики МГТУ</a></li>
							<li class="j-marg-bottom-xs"><a href="https://bmstu.ru/mstu/works/pr-media/press-kit/">Для СМИ</a></li>	
							<li class="j-marg-bottom-xs"><a href="https://bmstu.ru/mstu/contacts/emergency">Телефоны горячей линии и экстренных служб</a></li>
						</ul></td>
						</tr></tbody></table>
					</div>

					<div class="col-sm-4 col-md-4">
						<table><tbody><tr>
						<td class="visible-lg 	"><h3 class="j-marg-sm j-bold" style="font-size: 18px;"><i>Образовательный центр</i></h3></td>
						<td class="visible-lg"><img src="./МГТУ им. Н. Э. Баумана_files/decor-brace.png"></td>
						<td class="j-padd-left-sm">
							<ul class="j-vlist j-smaller j-padd-0">
							<li class="j-marg-bottom-xs"><a href="http://edu.bmstu.ru/napravleniya-obucheniya/programmirovanie/" target="_blank">Курсы по программированию</a></li>
							<li class="j-marg-bottom-xs"><a href="http://edu.bmstu.ru/napravleniya-obucheniya/marketing-i-reklama/" target="_blank">Маркетинг и реклама</a></li>
							<li class="j-marg-bottom-xs"><a href="http://edu.bmstu.ru/napravleniya-obucheniya/internet-marketing/" target="_blank">Интернет-маркетинг</a></li>
							<li class="j-marg-bottom-xs"><a href="http://edu.bmstu.ru/" target="_blank">Анализ данных и машинное обучение</a></li>
							<li class="j-marg-bottom-xs"><a href="http://edu.bmstu.ru/napravleniya-obucheniya/dlya-polzovatelej-pk/microsoft-office-excel/" target="_blank">Microsoft Excel</a></li>
							<li class="j-marg-bottom-xs"><a href="http://edu.bmstu.ru/napravleniya-obucheniya/buxgalterskij-uchet/" target="_blank">Бухгалтерский учет и аудит</a></li>
						</ul></td>
						</tr></tbody></table>
					</div>

					<div class="col-sm-4 col-md-4">
						<table><tbody><tr>
						<td class="visible-lg 	"><h3 class="j-marg-sm j-bold" style="font-size: 18px;"><i>Информация</i></h3></td>
						<td class="visible-lg"><img src="./МГТУ им. Н. Э. Баумана_files/decor-brace.png"></td>
						<td class="j-padd-left-sm">
							<ul class="j-vlist j-smaller j-padd-0">
								<li class="j-marg-bottom-xs j-bold"><a class="j-color-red" href="https://bmstu.ru/sveden/">Сведения об образовательной организации</a></li>	
							</ul>
						</td>
						</tr></tbody></table>
						<div style="height: 200px;"></div>
					</div>

				</div>
			</div>
		</div>

	</div>
	<div class="visible-xs j-padd-bottom-md">
		<div class="container">
			<div class="row">
				<div class="col-xs-6">
					<table><tbody><tr>
					<td><h4 class="j-color-gray">Полезное <span class="glyphicon glyphicon-info-sign"></span></h4>
						<ul class="j-vlist j-padd-0">
							<li class="j-marg-bottom-xs"><a href="https://bmstu.ru/mstu/about/university/organization#faculties-and-branches">Факультеты, Кафедры, Филиалы</a></li>
							<!-- <li class="j-marg-bottom-xs"><a href="/mstu/about/university/divisions/">Подразделения и Управления</a></li> -->
							<li class="j-marg-bottom-xs"><a href="https://bmstu.ru/mstu/about/university/documents/">Нормативные документы, Бланки</a></li>

							<li class="j-marg-bottom-xs"><a href="https://bmstu.ru/abitur/general/qanda/">Часто задаваемые вопросы</a></li>
							<li class="j-marg-bottom-xs"><a href="https://bmstu.ru/mstu/about/organization#research-centers">Научные центры, ЦКП и УНУ</a></li>
				
							
							<!-- <li class="j-marg-bottom-xs"><a href="/mstu/about/university/bauman-glossary/">Бауманский глоссарий</a> </li> -->
							<!-- <li class="j-marg-bottom-xs"><a href="/mstu/about/university/virtual-tour/">Виртуальная прогулка</a></li> -->
							<li class="j-marg-bottom-xs"><a href="http://shop.bmstu.ru/" target="_blank">Магазин атрибутики МГТУ</a></li>
							
							<li class="j-marg-bottom-xs"><a href="https://bmstu.ru/mstu/works/pr-media/press-kit/">Для СМИ</a></li>	
							<li class="j-marg-bottom-xs"><a href="https://bmstu.ru/mstu/contacts/emergency">Телефоны горячей линии и экстренных служб</a></li>
						</ul></td>
					</tr></tbody></table>
				</div>
					
				<div class="col-xs-6">
					<table><tbody><tr>
					<td><h4 class="j-color-gray">Информация <span class="glyphicon glyphicon-question-sign"></span></h4>
							<ul class="j-vlist j-padd-0">	
							<li class="j-marg-bottom-xs j-bold"><a href="https://bmstu.ru/sveden">Сведения об образовательной организации</a></li>
							<li class="j-marg-bottom-xs j-bold"><a class="j-color-red" href="https://bmstu.ru/feedback/">Обращения граждан</a></li>
						</ul></td>
					</tr></tbody></table>
				</div>
			</div>
		</div>
	</div>
	
	<div class="j-bg-white">
		<div class="container">
			<div class="row j-padd-top-sm j-padd-bottom-sm">

								
				<div id="events" class="b-events-test col-md-4 col-md-push-4 j-padd-top-sm">
					<p class="j-marg-0 text-center"><a href="https://bmstu.ru/mstu/info/events/">Лента событий</a></p>
					
					<div class="j-padd-md b-newsitem">
					<center>
						<div class="j-ribbon small j-display-iblock">9 мая 2020</div>
					</center>
					<h4 class="j-color-darkolive j-bold text-center j-padd-sm"><a class="j-color-darkolive" href="https://bmstu.ru/mstu/news/news.html?newsid=6984">Акция «Флаги России. 9 мая» прошла в МГТУ им. Н.Э. Баумана</a></h4>
					<div class="row"><!--lg-4-->
						<div class="col-sm-12 col-md-12 col-lg-12 j-padd-bottom-sm"><a href="https://bmstu.ru/mstu/news/news.html?newsid=6984"><span class="btn-video-icon"><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/decor-btn-play-100.png"></span><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_6984(3).jpg"></a></div>
					</div>
					<p class="j-padd-top-sm"><a href="https://bmstu.ru/mstu/news/news.html?newsid=6984"><small>9 мая студенты МГТУ им. Н.Э. Баумана приняли участие в  акции «Флаги России. 9 мая»....</small></a></p>
					<!--<p class="text-center"><a href="/mstu/news/news.html?newsid=6984" class="btn btn-default btn-sm"><span class="j-padd-left j-padd-right">Читать</span></a></p>-->
					<div class="j-hidden "></div>
				</div>					
				</div>
				<hr class="visible-sm">
				<div id="articles" class="col-sm-6 col-md-4 col-md-pull-4 j-padd-top-sm">
					<p class="j-padd-left-md j-marg-0"><a href="https://bmstu.ru/mstu/info/articles/">Видео и Статьи</a> &nbsp;<img src="./МГТУ им. Н. Э. Баумана_files/icon-star.gif" alt="">&nbsp; <a href="https://bmstu.ru/mstu/info/interviews/">Интервью</a>&nbsp;<img src="./МГТУ им. Н. Э. Баумана_files/icon-star.gif" alt="">&nbsp; <a href="https://bmstu.ru/mstu/works/pr-media/press-kit/">«Инженер»</a></p>


										
					
					<div class="j-padd-md b-newsitem">
					
					<p class="hidden-xs j-date">14 мая 2020</p>
					
					<div class="row j-padd-top-sm j-padd-bottom-sm">
					
						
						<div class="col-md-12 j-padd-bottom-sm">
							<table><tbody><tr><td><img src="./МГТУ им. Н. Э. Баумана_files/icon-tv.png"> <a href="https://bmstu.ru/mstu/works/pr-media/#tv-studio">МГТУ-ТВ</a></td></tr></tbody></table>
							
						</div>
						<div class="col-md-10 j-padd-bottom-sm"><center><a href="https://bmstu.ru/#video-channel" data-toggle="modal" data-target="#video" class="btn-video"><span class="btn-video-icon"><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/decor-btn-play-100.png"></span><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_6991.jpg"></a></center></div>
						
						<div class="col-md-10 j-padd-bottom-sm"><h4 class="text-center j-textblock"><a class="j-color-darkolive j-bold" href="https://bmstu.ru/mstu/news/news.html?newsid=6991">Георгий Васильевич ЗИМИН (1912-1997). Проект «Лица Победы».</a></h4></div>
					</div>
					<div class="j-hidden b-video-iframe"><iframe src="./МГТУ им. Н. Э. Баумана_files/lTZbvY9Cnuk.html" width="560" height="315" frameborder="0" webkitallowfullscreen="" mozallowfullscreen="" allowfullscreen=""></iframe></div>
				</div><div class="j-padd-md b-newsitem">
					
					<p class="hidden-xs j-date">7 мая 2020</p>
					
					<div class="row j-padd-top-sm j-padd-bottom-sm">
					
						
						<div class="col-md-12 j-padd-bottom-sm">
							<table><tbody><tr><td><img src="./МГТУ им. Н. Э. Баумана_files/icon-tv.png"> <a href="https://bmstu.ru/mstu/works/pr-media/#tv-studio">МГТУ-ТВ</a></td></tr></tbody></table>
							
						</div>
						<div class="col-md-10 j-padd-bottom-sm"><center><a href="https://bmstu.ru/#video-channel" data-toggle="modal" data-target="#video" class="btn-video"><span class="btn-video-icon"><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/decor-btn-play-100.png"></span><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_6979.jpg"></a></center></div>
						
						<div class="col-md-10 j-padd-bottom-sm"><h4 class="text-center j-textblock"><a class="j-color-darkolive j-bold" href="https://bmstu.ru/mstu/news/news.html?newsid=6979">Степан Николаевич ПЕРЕКАЛЬСКИЙ (1898-1943). Проект «Лица Победы».</a></h4></div>
					</div>
					<div class="j-hidden b-video-iframe"><iframe src="./МГТУ им. Н. Э. Баумана_files/0PBDIHDGi5k.html" width="560" height="315" frameborder="0" webkitallowfullscreen="" mozallowfullscreen="" allowfullscreen=""></iframe></div>
				</div>					
					
					
				</div>				

				<div id="news" class="col-sm-6 col-md-4 j-padd-top-sm">
					<p class="j-padd-left-md j-marg-0 j-align-right"><a href="https://bmstu.ru/mstu/info/bauman-news/">Новости университета</a></p>
					<div class="j-padd-left-sm j-padd-top-md j-padd-bottom-md">
					

						
					<div class="row j-marg-bottom-md b-newsitem">			
					<div class="col-md-12"><p class="j-date">9 июня 2020</p></div>
					<div class="col-md-5"><p class="hidden-xs hidden-sm"><a href="https://bmstu.ru/mstu/news/news.html?newsid=7027"><span class="btn-video-icon"><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/decor-btn-play-100.png"></span><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_7027.jpg"></a></p></div><div class="col-md-7"><p class="j-bold"><a class="j-color-darkolive" href="https://bmstu.ru/mstu/news/news.html?newsid=7027">Вышел приказ ректора об организации работы Университета в связи с поэтапным снятием ограничений в условиях неблагоприятной эпидемиологической ситуации</a></p></div>
				</div><div class="row j-marg-bottom-md b-newsitem">			
					<div class="col-md-12"><p class="j-date">12 июня 2020</p></div>
					<div class="col-md-5"><p class="hidden-xs hidden-sm"><a href="https://bmstu.ru/#video-news" data-toggle="modal" data-target="#video" class="btn-video"><span class="btn-video-icon"><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/decor-btn-play-100.png"></span><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_7033.jpg"></a></p></div><div class="col-md-7"><p class="j-bold"><a class="j-color-darkolive" href="https://bmstu.ru/mstu/news/news.html?newsid=7033">Самый дешевый российский автомобиль на спорте: Феррари из Лады - родстер КРЫМ</a></p></div><div class="j-hidden b-video-iframe"><iframe src="./МГТУ им. Н. Э. Баумана_files/uLOsOowQjuw.html" width="560" height="315" frameborder="0" webkitallowfullscreen="" mozallowfullscreen="" allowfullscreen=""></iframe></div>
				</div><div class="row j-marg-bottom-md b-newsitem">			
					<div class="col-md-12"><p class="j-date">11 июня 2020</p></div>
					<div class="col-md-5"><p class="hidden-xs hidden-sm"><a href="https://bmstu.ru/mstu/news/news.html?newsid=7032"><span class="btn-video-icon"><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/decor-btn-play-100.png"></span><img class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_7032.png"></a></p></div><div class="col-md-7"><p class="j-bold"><a class="j-color-darkolive" href="https://bmstu.ru/mstu/news/news.html?newsid=7032">Инженер года: Дмитрий Денисов.</a></p></div>
				</div>						
						
					
					</div>
				</div>	

			</div>

		</div>
	</div>

	<div class="j-bg-beige">
		<div class="j-relative"><div class="crown-decor"></div></div>
		<div class="container">
			<div class="row j-marg2">
				<div class="col-md-12"><p class="j-marg-0 text-center"><a class="j-larger" href="https://bmstu.ru/mstu/info/adverts/">Ближайшие мероприятия, Объявления</a></p>
					</div>
			</div>
			<div class="row j-padd-bottom">
				
				<div class="col-md-12">
					<div class="" style="padding : 0 30px;">
						<div class="autoplay-1 slick-initialized slick-slider slick-dotted" role="toolbar"><button type="button" data-role="none" class="slick-prev slick-arrow" aria-label="Previous" role="button" style="display: block;">Previous</button>
							<div aria-live="polite" class="slick-list draggable"><div class="slick-track" role="listbox" style="opacity: 1; width: 4320px; transform: translate3d(-1080px, 0px, 0px);"><div class="slick-slide slick-cloned" data-slick-index="-4" aria-hidden="true" tabindex="-1" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>15 июня 2020</small></p><p class="j-marg2"><img style="max-height:144px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_7029.png"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=7029" tabindex="-1">Прием заявок на основную программу «Технопарка» 

</a></p></div></div><div class="slick-slide slick-cloned" data-slick-index="-3" aria-hidden="true" tabindex="-1" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>1 июня 2020</small></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=7009" tabindex="-1">МГТУ им. Н.Э. Баумана продолжает обучение столичных преподавателей на базе технопарка «Инжинириум»
</a></p></div></div><div class="slick-slide slick-cloned" data-slick-index="-2" aria-hidden="true" tabindex="-1" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>10 июня 2020</small></p><p class="j-marg2"><img style="max-height:144px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_6987.jpg"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=6987" tabindex="-1">Золотые имена высшей школы

</a></p></div></div><div class="slick-slide slick-cloned" data-slick-index="-1" aria-hidden="true" tabindex="-1" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>1 июня 2020</small></p><p class="j-marg2"><img style="max-height:144px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_7016(1).jpg"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=7016" tabindex="-1">Приглашаем в студенческий строительный отряд для осуществления ремонтных работ на территории МГТУ</a></p></div></div><div class="slick-slide slick-current slick-active" data-slick-index="0" aria-hidden="false" tabindex="-1" role="option" aria-describedby="slick-slide00" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>18 мая 2020</small></p><p class="j-marg2"><img style="max-height:144px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_6997(3).jpg"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=6997" tabindex="0">Старт летнего оздоровительного сезона 2020

</a></p></div></div><div class="slick-slide slick-active" data-slick-index="1" aria-hidden="false" tabindex="-1" role="option" aria-describedby="slick-slide01" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>21 июня 2020</small></p><p class="j-marg2"><img style="max-height:144px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_7022.png"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=7022" tabindex="0">Команда Арктики ждет тебя!</a></p></div></div><div class="slick-slide slick-active" data-slick-index="2" aria-hidden="false" tabindex="-1" role="option" aria-describedby="slick-slide02" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>6 июня 2020</small></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=7024" tabindex="0">Набор офицеров, сержантов запаса в 2020г.

.</a></p></div></div><div class="slick-slide slick-active" data-slick-index="3" aria-hidden="false" tabindex="-1" role="option" aria-describedby="slick-slide03" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>6 июня 2020</small></p><p class="j-marg2"><img style="max-height:144px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_7025.png"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=7025" tabindex="0">Информация о приемной кампании военного учебного центра в 2020г</a></p></div></div><div class="slick-slide" data-slick-index="4" aria-hidden="true" tabindex="-1" role="option" aria-describedby="slick-slide04" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>15 июня 2020</small></p><p class="j-marg2"><img style="max-height:144px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_7029.png"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=7029" tabindex="-1">Прием заявок на основную программу «Технопарка» 

</a></p></div></div><div class="slick-slide" data-slick-index="5" aria-hidden="true" tabindex="-1" role="option" aria-describedby="slick-slide05" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>1 июня 2020</small></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=7009" tabindex="-1">МГТУ им. Н.Э. Баумана продолжает обучение столичных преподавателей на базе технопарка «Инжинириум»
</a></p></div></div><div class="slick-slide" data-slick-index="6" aria-hidden="true" tabindex="-1" role="option" aria-describedby="slick-slide06" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>10 июня 2020</small></p><p class="j-marg2"><img style="max-height:144px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_6987.jpg"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=6987" tabindex="-1">Золотые имена высшей школы

</a></p></div></div><div class="slick-slide" data-slick-index="7" aria-hidden="true" tabindex="-1" role="option" aria-describedby="slick-slide07" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>1 июня 2020</small></p><p class="j-marg2"><img style="max-height:144px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_7016(1).jpg"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=7016" tabindex="-1">Приглашаем в студенческий строительный отряд для осуществления ремонтных работ на территории МГТУ</a></p></div></div><div class="slick-slide slick-cloned" data-slick-index="8" aria-hidden="true" tabindex="-1" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>18 мая 2020</small></p><p class="j-marg2"><img style="max-height:144px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_6997(3).jpg"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=6997" tabindex="-1">Старт летнего оздоровительного сезона 2020

</a></p></div></div><div class="slick-slide slick-cloned" data-slick-index="9" aria-hidden="true" tabindex="-1" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>21 июня 2020</small></p><p class="j-marg2"><img style="max-height:144px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_7022.png"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=7022" tabindex="-1">Команда Арктики ждет тебя!</a></p></div></div><div class="slick-slide slick-cloned" data-slick-index="10" aria-hidden="true" tabindex="-1" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>6 июня 2020</small></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=7024" tabindex="-1">Набор офицеров, сержантов запаса в 2020г.

.</a></p></div></div><div class="slick-slide slick-cloned" data-slick-index="11" aria-hidden="true" tabindex="-1" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>6 июня 2020</small></p><p class="j-marg2"><img style="max-height:144px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_7025.png"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=7025" tabindex="-1">Информация о приемной кампании военного учебного центра в 2020г</a></p></div></div></div></div>						
						<button type="button" data-role="none" class="slick-next slick-arrow" aria-label="Next" role="button" style="display: block;">Next</button><ul class="slick-dots" style="display: block;" role="tablist"><li class="slick-active" aria-hidden="false" role="presentation" aria-selected="true" aria-controls="navigation00" id="slick-slide00"><button type="button" data-role="none" role="button" tabindex="0">1</button></li><li aria-hidden="true" role="presentation" aria-selected="false" aria-controls="navigation01" id="slick-slide01" class=""><button type="button" data-role="none" role="button" tabindex="0">2</button></li><li aria-hidden="true" role="presentation" aria-selected="false" aria-controls="navigation02" id="slick-slide02" class=""><button type="button" data-role="none" role="button" tabindex="0">3</button></li><li aria-hidden="true" role="presentation" aria-selected="false" aria-controls="navigation03" id="slick-slide03" class=""><button type="button" data-role="none" role="button" tabindex="0">4</button></li><li aria-hidden="true" role="presentation" aria-selected="false" aria-controls="navigation04" id="slick-slide04" class=""><button type="button" data-role="none" role="button" tabindex="0">5</button></li><li aria-hidden="true" role="presentation" aria-selected="false" aria-controls="navigation05" id="slick-slide05" class=""><button type="button" data-role="none" role="button" tabindex="0">6</button></li><li aria-hidden="true" role="presentation" aria-selected="false" aria-controls="navigation06" id="slick-slide06" class=""><button type="button" data-role="none" role="button" tabindex="0">7</button></li><li aria-hidden="true" role="presentation" aria-selected="false" aria-controls="navigation07" id="slick-slide07" class=""><button type="button" data-role="none" role="button" tabindex="0">8</button></li></ul></div>
					</div>
				
				</div>
				
					
					


			</div>
		</div>

	</div>	

	<div class="j-bg-white">
		<div class="container">
			<div class="row j-marg2">
				<div class="col-md-12"><p class="j-marg-0 text-center"><a class="j-larger" href="https://bmstu.ru/mstu/info/upcoming-events/">Текущие мероприятия</a></p>
					</div>
			</div>
			<div class="row j-padd-bottom">
				
				<div class="col-md-12">
					<div class="" style="padding : 0 30px;">
						<div class="autoplay-2 slick-initialized slick-slider">
							<div aria-live="polite" class="slick-list draggable"><div class="slick-track" role="listbox" style="opacity: 1; width: 1080px; transform: translate3d(0px, 0px, 0px);"><div class="slick-slide slick-current slick-active" data-slick-index="0" aria-hidden="false" tabindex="-1" role="option" aria-describedby="slick-slide10" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>1 января — 31 декабря</small></p><p class="j-marg2"><img style="max-height:144px;max-width:190px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_6548(1).jpg"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=6548" tabindex="0">VII Международный суворовский патриотический фестиваль искусств</a></p></div></div><div class="slick-slide slick-active" data-slick-index="1" aria-hidden="false" tabindex="-1" role="option" aria-describedby="slick-slide11" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>2 марта — 30 ноября</small></p><p class="j-marg2"><img style="max-height:144px;max-width:190px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_6826.jpg"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=6826" tabindex="0">Всероссийский марафон ценностей здорового образа жизни «Заряжайся на здоровье»</a></p></div></div><div class="slick-slide slick-active" data-slick-index="2" aria-hidden="false" tabindex="-1" role="option" aria-describedby="slick-slide12" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>31 мая — 6 октября</small></p><p class="j-marg2"><img style="max-height:144px;max-width:190px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_6954.jpg"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=6954" tabindex="0">III Всероссийская научно-практическая конференция по системам управления жизненным циклом продукции</a></p></div></div><div class="slick-slide slick-active" data-slick-index="3" aria-hidden="false" tabindex="-1" role="option" aria-describedby="slick-slide13" style="width: 270px;"><div style="width:80%;"><p class="j-date j-color-red j-bold"><small>1 мая — 30 июля</small></p><p class="j-marg2"><img style="max-height:144px;max-width:190px;" class="img-responsive" src="./МГТУ им. Н. Э. Баумана_files/img_6969.jpg"></p><p><a href="https://bmstu.ru/mstu/news/news.html?newsid=6969" tabindex="0">Подведение итогов выполнения Коллективного Договора 2017-2020</a></p></div></div></div></div>						
						</div>
					</div>
				
				</div>				
				

			</div>
		</div>
	</div>		
	


	<div class="j-bg-beige">
		<div class="j-relative"><div class="crown-decor"></div></div>
		<div class="container">
			<div class="row j-marg2">
				<div class="col-md-12"><p class="j-marg-0 text-center"><a target="_blank" href="https://bus.gov.ru/">Официальный сайт для размещения информации
о государственных (муниципальных) учреждениях и их независимой оценке</a></p>
					</div>
			</div>
		</div>
	</div>	


	
	<div id="footer">
		<div class="container">
			<div class="hidden-xs row j-padd-top ">
				

				<div class="col-sm-5 col-md-4 col-md-push-8 j-smaller">

					<ul class="list-inline">
						<!-- <li class="j-padd-0"><a href="https://www.facebook.com/baumanski?ref=aymt_homepage_panel"><img src="/b-twig/assets/img/icon-facebook.png"></a>&nbsp;</li> -->
						<li class="j-padd-0"><a href="https://www.facebook.com/bmstu1830/"><img src="./МГТУ им. Н. Э. Баумана_files/icon-facebook.png"></a>&nbsp;</li>
						<li class="j-padd-0"><a href="http://www.youtube.com/user/TVbaumanka"><img src="./МГТУ им. Н. Э. Баумана_files/icon-youtube.png"></a>&nbsp;</li>
						<li class="j-padd-0"><a href="https://twitter.com/bmstu1830"><img src="./МГТУ им. Н. Э. Баумана_files/icon-twitter.png"></a>&nbsp;</li>
						<!-- <li class="j-padd-0"><a href="http://vk.com/pr.bmstu"><img src="/b-twig/assets/img/icon-vk.png"></a>&nbsp;</li> -->
						<li class="j-padd-0"><a href="https://vk.com/bmstu1830"><img src="./МГТУ им. Н. Э. Баумана_files/icon-vk.png"></a>&nbsp;</li>
						<li class="j-padd-0"><a href="https://www.instagram.com/bmstu1830/"><img alt="" src="./МГТУ им. Н. Э. Баумана_files/icon-instagram.png"></a>&nbsp;</li>

					</ul>
					<ul class="j-vlist j-padd-0 j-marg-top-sm">
						<li>&nbsp;</li>
						<li><a href="https://bmstu.ru/mstu/antiterror/">Антитеррор</a></li>
						<li><a href="https://bmstu.ru/mstu/about/anticorr/">Противодействие коррупции</a></li>
						<li><a href="https://bmstu.ru/mstu/support/contact-us/">Обратная связь</a></li>
						<li><a href="https://bmstu.ru/mstu/sitemap/">Карта сайта</a></li>
						<li>&nbsp;</li>
						<li>Разработано и оформлено отделом <br><a href="https://bmstu.ru/mstu/support/contact-us/">Поддержки интернет-ресурсов и баз <br>данных УИМП</a> МГТУ&nbsp;им.&nbsp;Н.&nbsp;Э.&nbsp;Баумана</li>
						<li>&nbsp;</li>
					</ul>
				</div>					
				
				
				<div class="col-sm-1 col-md-4 j-italic j-color-darkolive" style="font-family:Georgia;">
					<div class="hidden-sm j-padd-right-sm">
						<p class="j-bold">Ракетный колледж на Яузе</p>
						<p>— так иногда называют наш университет. <br>МВТУ — пионер в решении задач авиации, создании ракетной техники, ядерной энергетики и радиоэлектроники</p>
					</div>
				</div>
				


				
				<div class="col-sm-6 col-md-4 col-md-pull-8 j-smaller">
					
					<div class="contacts">
						<div class="logo-grayscale"></div>
						<div class="contacts-body j-color-gray j-padd-right">
							<p class="j-bold">105005, Москва, 2-я&nbsp;Бауманская&nbsp;ул., д.&nbsp;5, стр.&nbsp;1</p>
							<table><tbody>
								<tr><td class="j-padd-right-sm">Телефон</td><td class="j-bold">(499) 263-6391</td></tr>
								<tr><td class="j-padd-right-sm">Факс</td><td class="j-bold">(499) 267-4844</td></tr>
								<tr><td class="j-padd-right-sm">Эл.&nbsp;почта</td><td class="j-bold">bauman@bmstu.ru</td></tr>
							</tbody></table>
							<p class="j-bold j-marg-top j-color-red">Приемная комиссия:</p>
							<table><tbody>
								<tr><td class="j-padd-right-sm">Телефон</td><td class="j-bold">(499) 263-6541</td></tr>
								<tr><td class="j-padd-right-sm">Эл.&nbsp;почта</td><td class="j-bold">abiturient@bmstu.ru</td></tr>
								<tr><td class="j-padd-right-sm j-bold" colspan="2"><br></td></tr>
							</tbody></table>
							<p class="j-padd-top-sm j-bold"><a href="https://bmstu.ru/mstu/contacts/map/">Схема проезда</a></p>
						</div>
					</div>

				</div>
				
			</div>
			<div class="visible-xs row j-padd-top">
				<div class="col-xs-12">
					
					
					<div class="contacts">
						<div class="logo-grayscale"></div>
						<div class="contacts-body j-color-gray j-padd-right">
							<ul class="list-inline">
								<!-- <li class="j-padd-0"><a href="https://www.facebook.com/baumanski?ref=aymt_homepage_panel"><img src="/b-twig/assets/img/icon-facebook.png"></a>&nbsp;</li> -->
								<li class="j-padd-0"><a href="https://www.facebook.com/bmstu1830/"><img src="./МГТУ им. Н. Э. Баумана_files/icon-facebook.png"></a>&nbsp;</li>
								<li class="j-padd-0"><a href="http://www.youtube.com/user/TVbaumanka"><img src="./МГТУ им. Н. Э. Баумана_files/icon-youtube.png"></a>&nbsp;</li>
								<li class="j-padd-0"><a href="https://twitter.com/bmstu1830"><img src="./МГТУ им. Н. Э. Баумана_files/icon-twitter.png"></a>&nbsp;</li>
								<!-- <li class="j-padd-0"><a href="http://vk.com/pr.bmstu"><img src="/b-twig/assets/img/icon-vk.png"></a>&nbsp;</li> -->
								<li class="j-padd-0"><a href="https://vk.com/bmstu1830"><img src="./МГТУ им. Н. Э. Баумана_files/icon-vk.png"></a>&nbsp;</li>
								<li class="j-padd-0"><a href="https://www.instagram.com/bmstu1830/"><img alt="" src="./МГТУ им. Н. Э. Баумана_files/icon-instagram.png"></a>&nbsp;</li>
								<!-- VK47662 -->
							</ul>						
							<p class="j-bold">105005, Москва, 2-я&nbsp;Бауманская&nbsp;ул., д.&nbsp;5, стр.&nbsp;1</p>
							<table><tbody>
								<tr><td class="j-padd-right-sm">Телефон</td><td class="j-bold">(499) 263-6391</td></tr>
								<tr><td class="j-padd-right-sm">Факс</td><td class="j-bold">(499) 267-4844</td></tr>
								<tr><td class="j-padd-right-sm">Эл.&nbsp;почта</td><td class="j-bold">bauman@bmstu.ru</td></tr>
							</tbody></table>
							<p class="j-bold j-marg-top j-color-red">Приемная комиссия:</p>
							<table><tbody>
								<tr><td class="j-padd-right-sm">Телефон</td><td class="j-bold">(499) 263-6541</td></tr>
								<tr><td class="j-padd-right-sm">Эл.&nbsp;почта</td><td class="j-bold">abiturient@bmstu.ru</td></tr>
								<tr><td class="j-padd-right-sm j-bold" colspan="2"><br></td></tr>
							</tbody></table>
							<p class="j-padd-top-sm j-bold"><a class="btn btn-default" href="https://bmstu.ru/mstu/contacts/map/">Схема проезда</a></p>
						</div>
					</div>

				</div>
			</div>
		</div>
	</div>
	
	
</div>




<div class="j-hidden">


	<div id="ie-alert" class="j-padd-top-sm j-padd-bottom alert-dismissable">
		<button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>
		<div class="container">
			<div class="row">
				<div class="col-md-12">			
					<h3><span>Ваш браузер — Internet Explorer</span></h3>
					<p><b>К сожалению, этот браузер уже устарел</b>: он уже не поддерживает новые веб-технологии и не соответствует современным веб-стандартам, поэтому некоторые элементы на странице могут отображаться некорректно. В этой связи, рекомендуется обновить Ваш браузер до <b>последней версии</b> или использовать альтернативные браузеры бесплатно, такие как <b>Google Chrome</b>, <b>Mozilla Firefox</b>, <b>Yandex</b>, <b>Opera</b></p>

				</div>
			</div>
		</div>
	</div>
	
	
</div>

<div id="eye-impaired" class="modal fade bs-example-modal-sm" tabindex="-1" role="dialog" aria-labelledby="mySmallModalLabel">
	<div class="modal-dialog modal-sm">
		<div class="modal-content">
			<div class="modal-header">
			<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">×</span></button>
			<h3 class="modal-title j-padd-lg" id="myModalLabel">Для людей с ограниченными возможностями</h3>
			</div>
			<div class="modal-body"><p class="j-padd j-larger" style="font-size: 20px !important;">
			  Для комфортного просмотра вы можете воспользоваться масштабированием страницы — в активном окне браузера нажмите несколько раз на клавиатуре сочетание клавиш <code>Ctrl</code> и <code>+</code> для увеличения размеров шрифта и элементов сайта до необходимого вам размера</p>
			</div>
			<div class="modal-footer">
			<button type="button" class="btn btn-primary" data-dismiss="modal">Закрыть</button>
			</div>
		</div>
	</div>  
</div>
<div id="video" class="modal fade bs-example-modal-sm" tabindex="-1" role="dialog" aria-labelledby="">
	<div class="modal-dialog modal-sm">
		<div class="modal-content modal-content-invert">
			<div class="modal-header"><button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">×</span></button></div>
			<div class="modal-body">
				<div class="b-video"></div>
			</div>
			<div class="modal-footer">
			<button type="button" class="btn btn-default btn-sm" data-dismiss="modal">Закрыть</button>
			</div>
		</div>
	</div>  
</div>
<!--
<script type="text/javascript" src="https://code.jquery.com/jquery.js"></script>
-->
<script async="" type="text/javascript" src="./МГТУ им. Н. Э. Баумана_files/cnt.js"></script><script type="text/javascript" src="./МГТУ им. Н. Э. Баумана_files/jquery-1.7.1.min.js"></script>
<script type="text/javascript" src="./МГТУ им. Н. Э. Баумана_files/bootstrap.min.js"></script>
<script type="text/javascript" src="./МГТУ им. Н. Э. Баумана_files/holder.js"></script>
<script type="text/javascript" src="./МГТУ им. Н. Э. Баумана_files/responsive-carousel-controls.js"></script>
<script type="text/javascript" src="./МГТУ им. Н. Э. Баумана_files/bowser.min.js"></script>
<script type="text/javascript" src="./МГТУ им. Н. Э. Баумана_files/init.js"></script>
<script type="text/javascript" src="./МГТУ им. Н. Э. Баумана_files/slick.min.js"></script>
<script src="./МГТУ им. Н. Э. Баумана_files/snowfall.jquery.min.js"></script>

<script type="text/javascript">
$(document).ready(function(){

	$('.autoplay-1').slick({
		slidesToShow: 4,
		slidesToScroll: 1,
		autoplay: true,
		autoplaySpeed: 3000,
		dots: true,
		responsive: [
			{
				breakpoint: 992,
				settings: {
					slidesToShow: 3,
					slidesToScroll: 3,
					infinite: true,
					dots: true
				}
			},
			{
				breakpoint: 768,
				settings: {
					slidesToShow: 2,
					slidesToScroll: 2
				}
			},
			{
				breakpoint: 480,
				settings: {
					slidesToShow: 1,
					slidesToScroll: 1
				}
			}
		]
	});	
	$('.autoplay-2').slick({
		slidesToShow: 4,
		slidesToScroll: 1,
		autoplay: true,
		dots: true,
		autoplaySpeed: 3000,
		responsive: [
			{
				breakpoint: 992,
				settings: {
					slidesToShow: 3,
					slidesToScroll: 3,
					infinite: true,
					dots: true
				}
			},
			{
				breakpoint: 768,
				settings: {
					slidesToShow: 2,
					slidesToScroll: 2
				}
			},
			{
				breakpoint: 480,
				settings: {
					slidesToShow: 1,
					slidesToScroll: 1
				}
			}
		]
	});	
});

</script>
<script type="text/javascript">
       (function(d, t, p) {
           var j = d.createElement(t); j.async = true; j.type = "text/javascript";
           j.src = ("https:" == p ? "https:" : "http:") + "//stat.sputnik.ru/cnt.js";
           var s = d.getElementsByTagName(t)[0]; s.parentNode.insertBefore(j, s);
       })(document, "script", document.location.protocol);
</script>

<script src="./МГТУ им. Н. Э. Баумана_files/urchin.js" type="text/javascript"></script>
<script type="text/javascript">
_uacct = "UA-3048687-1";
urchinTracker();
</script>

<!-- VK730 -->



</body></html>")
}

func main() {
	worker := flag.Int("worker", 1, "worker count")
	flag.Parse()
	runtime.GOMAXPROCS(*worker)

	http.HandleFunc("/", hello)
	http.ListenAndServe(":5000", nil)
}
