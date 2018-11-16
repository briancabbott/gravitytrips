
package nasa_exoplanets_archive_damnd

/*
# COLUMN pl_hostname:    Host Name
# COLUMN pl_letter:      Planet Letter
# COLUMN pl_name:        Planet Name
# COLUMN pl_discmethod:  Discovery Method
# COLUMN pl_pnum:        Number of Planets in System
# COLUMN pl_orbper:      Orbital Period [days]
# COLUMN pl_orbpererr1:  Orbital Period Upper Unc. [days]
# COLUMN pl_orbpererr2:  Orbital Period Lower Unc. [days]
# COLUMN pl_orbperlim:   Orbital Period Limit Flag
# COLUMN pl_orbsmax:     Orbit Semi-Major Axis [AU])
# COLUMN pl_orbsmaxerr1: Orbit Semi-Major Axis Upper Unc. [AU]
# COLUMN pl_orbsmaxerr2: Orbit Semi-Major Axis Lower Unc. [AU]
# COLUMN pl_orbsmaxlim:  Orbit Semi-Major Axis Limit Flag
# COLUMN pl_orbeccen:    Eccentricity
# COLUMN pl_orbeccenerr1: Eccentricity Upper Unc.
# COLUMN pl_orbeccenerr2: Eccentricity Lower Unc.
# COLUMN pl_orbeccenlim: Eccentricity Limit Flag
# COLUMN pl_orbincl:     Inclination [deg]
# COLUMN pl_orbinclerr1: Inclination Upper Unc. [deg]
# COLUMN pl_orbinclerr2: Inclination Lower Unc. [deg]
# COLUMN pl_orbincllim:  Inclination Limit Flag
# COLUMN pl_bmassj:      Planet Mass or M*sin(i) [Jupiter mass]
# COLUMN pl_bmassjerr1:  Planet Mass or M*sin(i) Upper Unc. [Jupiter mass]
# COLUMN pl_bmassjerr2:  Planet Mass or M*sin(i) Lower Unc. [Jupiter mass]
# COLUMN pl_bmassjlim:   Planet Mass or M*sin(i) Limit Flag
# COLUMN pl_bmassprov:   Planet Mass or M*sin(i) Provenance
# COLUMN pl_radj:        Planet Radius [Jupiter radii]
# COLUMN pl_radjerr1:    Planet Radius Upper Unc. [Jupiter radii]
# COLUMN pl_radjerr2:    Planet Radius Lower Unc. [Jupiter radii]
# COLUMN pl_radjlim:     Planet Radius Limit Flag
# COLUMN pl_dens:        Planet Density [g/cm**3]
# COLUMN pl_denserr1:    Planet Density Upper Unc. [g/cm**3]
# COLUMN pl_denserr2:    Planet Density Lower Unc. [g/cm**3]
# COLUMN pl_denslim:     Planet Density Limit Flag
# COLUMN pl_ttvflag:     TTV Flag
# COLUMN pl_kepflag:     Kepler Field Flag
# COLUMN pl_k2flag:      K2 Mission Flag
# COLUMN pl_nnotes:      Number of Notes
# COLUMN ra_str:         RA [sexagesimal]
# COLUMN ra:             RA [decimal degrees]
# COLUMN dec_str:        Dec [sexagesimal]
# COLUMN dec:            Dec [decimal degrees]
# COLUMN st_dist:        Distance [pc]
# COLUMN st_disterr1:    Distance Upper Unc. [pc]
# COLUMN st_disterr2:    Distance Lower Unc. [pc]
# COLUMN st_distlim:     Distance Limit Flag
# COLUMN st_optmag:      Optical Magnitude [mag]
# COLUMN st_optmagerr:   Optical Magnitude Unc. [mag]
# COLUMN st_optmaglim:   Optical Magnitude Limit Flag
# COLUMN st_optband:     Optical Magnitude Band
# COLUMN gaia_gmag:      G-band (Gaia) [mag]
# COLUMN gaia_gmagerr:   G-band (Gaia) Unc. [mag]
# COLUMN gaia_gmaglim:   G-band (Gaia) Limit Flag
# COLUMN st_teff:        Effective Temperature [K]
# COLUMN st_tefferr1:    Effective Temperature Upper Unc. [K]
# COLUMN st_tefferr2:    Effective Temperature Lower Unc. [K]
# COLUMN st_tefflim:     Effective Temperature Limit Flag
# COLUMN st_mass:        Stellar Mass [Solar mass]
# COLUMN st_masserr1:    Stellar Mass Upper Unc. [Solar mass]
# COLUMN st_masserr2:    Stellar Mass Lower Unc. [Solar mass]
# COLUMN st_masslim:     Stellar Mass Limit Flag
# COLUMN st_rad:         Stellar Radius [Solar radii]
# COLUMN st_raderr1:     Stellar Radius Upper Unc. [Solar radii]
# COLUMN st_raderr2:     Stellar Radius Lower Unc. [Solar radii]
# COLUMN st_radlim:      Stellar Radius Limit Flag
# COLUMN rowupdate:      Date of Last Update
# COLUMN pl_tranflag:    Planet Transit Flag
# COLUMN pl_rvflag:      Planet RV Flag
# COLUMN pl_imgflag:     Planet Imaging Flag
# COLUMN pl_astflag:     Planet Astrometry Flag
# COLUMN pl_omflag:      Planet Orbital Modulation Flag
# COLUMN pl_cbflag:      Planet Circumbinary Flag
# COLUMN pl_angsep:      Calculated Angular Separation [mas]
# COLUMN pl_angseperr1:  Calculated Angular Separation Upper Unc. [mas]
# COLUMN pl_angseperr2:  Calculated Angular Separation Lower Unc. [mas]
# COLUMN pl_orbtper:     Time of Periastron [days]
# COLUMN pl_orbtpererr1: Time of Periastron Upper Unc. [days]
# COLUMN pl_orbtpererr2: Time of Periastron Lower Unc. [days]
# COLUMN pl_orbtperlim:  Time of Periastron Limit Flag
# COLUMN pl_orblper:     Long. of Periastron [deg]
# COLUMN pl_orblpererr1: Long. of Periastron Upper Unc. [deg]
# COLUMN pl_orblpererr2: Long. of Periastron Lower Unc. [deg]
# COLUMN pl_orblperlim:  Long. of Periastron Limit Flag
# COLUMN pl_rvamp:       Radial Velocity Amplitude [m/s]
# COLUMN pl_rvamperr1:   Radial Velocity Amplitude Upper Unc. [m/s]
# COLUMN pl_rvamperr2:   Radial Velocity Amplitude Lower Unc. [m/s]
# COLUMN pl_rvamplim:    Radial Velocity Amplitude Limit Flag
# COLUMN pl_eqt:         Equilibrium Temperature [K]
# COLUMN pl_eqterr1:     Equilibrium Temperature Upper Unc. [K]
# COLUMN pl_eqterr2:     Equilibrium Temperature Lower Unc. [K]
# COLUMN pl_eqtlim:      Equilibrium Temperature Limit Flag
# COLUMN pl_insol:       Insolation Flux [Earth flux]
# COLUMN pl_insolerr1:   Insolation Flux Upper Unc. [Earth flux]
# COLUMN pl_insolerr2:   Insolation Flux Lower Unc. [Earth flux]
# COLUMN pl_insollim:    Insolation Flux Limit Flag
# COLUMN pl_massj:       Planet Mass [Jupiter mass]
# COLUMN pl_massjerr1:   Planet Mass Upper Unc. [Jupiter mass]
# COLUMN pl_massjerr2:   Planet Mass Lower Unc. [Jupiter mass]
# COLUMN pl_massjlim:    Planet Mass Limit Flag
# COLUMN pl_msinij:      Planet M*sin(i) [Jupiter mass]
# COLUMN pl_msinijerr1:  Planet M*sin(i) Upper Unc. [Jupiter mass]
# COLUMN pl_msinijerr2:  Planet M*sin(i) Lower Unc. [Jupiter mass]
# COLUMN pl_msinijlim:   Planet M*sin(i) Limit Flag
# COLUMN pl_masse:       Planet Mass [Earth mass]
# COLUMN pl_masseerr1:   Planet Mass Upper Unc. [Earth mass]
# COLUMN pl_masseerr2:   Planet Mass Lower Unc. [Earth mass]
# COLUMN pl_masselim:    Planet Mass Limit Flag
# COLUMN pl_msinie:      Planet M*sin(i) [Earth mass]
# COLUMN pl_msinieerr1:  Planet M*sin(i) Upper Unc. [Earth mass]
# COLUMN pl_msinieerr2:  Planet M*sin(i) Lower Unc. [Earth mass]
# COLUMN pl_msinielim:   Planet M*sin(i) Limit Flag
# COLUMN pl_bmasse:      Planet Mass or M*sin(i) [Earth mass]
# COLUMN pl_bmasseerr1:  Planet Mass or M*sin(i) Upper Unc. [Earth mass]
# COLUMN pl_bmasseerr2:  Planet Mass or M*sin(i) Lower Unc. [Earth mass]
# COLUMN pl_bmasselim:   Planet Mass or M*sin(i) Limit Flag
# COLUMN pl_rade:        Planet Radius [Earth radii]
# COLUMN pl_radeerr1:    Planet Radius Upper Unc. [Earth radii]
# COLUMN pl_radeerr2:    Planet Radius Lower Unc. [Earth radii]
# COLUMN pl_radelim:     Planet Radius Limit Flag
# COLUMN pl_rads:        Planet Radius [Solar radii]
# COLUMN pl_radserr1:    Planet Radius Upper Unc. [Solar radii]
# COLUMN pl_radserr2:    Planet Radius Lower Unc. [Solar radii]
# COLUMN pl_radslim:     Planet Radius Limit Flag
# COLUMN pl_trandep:     Transit Depth [percent]
# COLUMN pl_trandeperr1: Transit Depth Upper Unc. [percent]
# COLUMN pl_trandeperr2: Transit Depth Lower Unc. [percent]
# COLUMN pl_trandeplim:  Transit Depth Limit Flag
# COLUMN pl_trandur:     Transit Duration [days]
# COLUMN pl_trandurerr1: Transit Duration Upper Unc. [days]
# COLUMN pl_trandurerr2: Transit Duration Lower Unc. [days]
# COLUMN pl_trandurlim:  Transit Duration Limit Flag
# COLUMN pl_tranmid:     Transit Midpoint [days]
# COLUMN pl_tranmiderr1: Transit Midpoint Upper Unc. [days]
# COLUMN pl_tranmiderr2: Transit Midpoint Lower Unc. [days]
# COLUMN pl_tranmidlim:  Transit Midpoint Limit Flag
# COLUMN pl_tsystemref:  Time System Reference
# COLUMN pl_imppar:      Impact Parameter
# COLUMN pl_impparerr1:  Impact Parameter Upper Unc.
# COLUMN pl_impparerr2:  Impact Parameter Lower Unc.
# COLUMN pl_impparlim:   Impact Parameter Limit Flag
# COLUMN pl_occdep:      Occultation Depth [percentage]
# COLUMN pl_occdeperr1:  Occultation Depth Upper Unc. [percentage]
# COLUMN pl_occdeperr2:  Occultation Depth Lower Unc. [percentage]
# COLUMN pl_occdeplim:   Occultation Depth Limit Flag
# COLUMN pl_ratdor:      Ratio of Distance to Stellar Radius
# COLUMN pl_ratdorerr1:  Ratio of Distance to Stellar Radius Upper Unc.
# COLUMN pl_ratdorerr2:  Ratio of Distance to Stellar Radius Lower Unc.
# COLUMN pl_ratdorlim:   Ratio of Distance to Stellar Radius Limit Flag
# COLUMN pl_ratror:      Ratio of Planet to Stellar Radius
# COLUMN pl_ratrorerr1:  Ratio of Planet to Stellar Radius Upper Unc.
# COLUMN pl_ratrorerr2:  Ratio of Planet to Stellar Radius Lower Unc.
# COLUMN pl_ratrorlim:   Ratio of Planet to Stellar Radius Limit Flag
# COLUMN pl_def_reflink: Default Reference
# COLUMN pl_disc:        Year of Discovery
# COLUMN pl_disc_reflink: Discovery Reference
# COLUMN pl_locale:      Discovery Locale
# COLUMN pl_facility:    Discovery Facility
# COLUMN pl_telescope:   Discovery Telescope
# COLUMN pl_instrument:  Discovery Instrument
# COLUMN pl_status:      Status
# COLUMN pl_mnum:        Number of Moons in System
# COLUMN pl_st_npar:     Number of Stellar and Planet Parameters
# COLUMN pl_st_nref:     Number of Stellar and Planet References
# COLUMN pl_pelink:      Link to Exoplanet Encyclopaedia
# COLUMN pl_edelink:     Link to Exoplanet Data Explorer
# COLUMN pl_publ_date:   Publication Date
# COLUMN hd_name:        HD Name
# COLUMN hip_name:       HIP Name
# COLUMN st_rah:         RA [hrs]
# COLUMN st_glon:        Galactic Longitude [deg]
# COLUMN st_glat:        Galactic Latitude [deg]
# COLUMN st_elon:        Ecliptic Longitude [deg]
# COLUMN st_elat:        Ecliptic Latitude [deg]
# COLUMN st_plx:         Parallax [mas]
# COLUMN st_plxerr1:     Parallax Upper Unc. [mas]
# COLUMN st_plxerr2:     Parallax Lower Unc. [mas]
# COLUMN st_plxlim:      Parallax Limit Flag
# COLUMN gaia_plx:       Gaia Parallax [mas]
# COLUMN gaia_plxerr1:   Gaia Parallax Upper Unc. [mas]
# COLUMN gaia_plxerr2:   Gaia Parallax Lower Unc. [mas]
# COLUMN gaia_plxlim:    Gaia Parallax Limit Flag
# COLUMN gaia_dist:      Gaia Distance [pc]
# COLUMN gaia_disterr1:  Gaia Distance Upper Unc. [pc]
# COLUMN gaia_disterr2:  Gaia Distance Lower Unc. [pc]
# COLUMN gaia_distlim:   Gaia Distance Limit Flag
# COLUMN st_pmra:        Proper Motion (RA) [mas/yr]
# COLUMN st_pmraerr:     Proper Motion (RA) Unc. [mas/yr]
# COLUMN st_pmralim:     Proper Motion (RA) Limit Flag
# COLUMN st_pmdec:       Proper Motion (Dec) [mas/yr]
# COLUMN st_pmdecerr:    Proper Motion (Dec) Unc. [mas/yr]
# COLUMN st_pmdeclim:    Proper Motion (Dec) Limit Flag
# COLUMN st_pm:          Total Proper Motion [mas/yr]
# COLUMN st_pmerr:       Total Proper Motion Unc. [mas/yr]
# COLUMN st_pmlim:       Total Proper Motion Limit Flag
# COLUMN gaia_pmra:      Gaia Proper Motion (RA) [mas/yr]
# COLUMN gaia_pmraerr:   Gaia Proper Motion (RA) Unc. [mas/yr]
# COLUMN gaia_pmralim:   Gaia Proper Motion (RA) Limit Flag
# COLUMN gaia_pmdec:     Gaia Proper Motion (Dec) [mas/yr]
# COLUMN gaia_pmdecerr:  Gaia Proper Motion (Dec) Unc. [mas/yr]
# COLUMN gaia_pmdeclim:  Gaia Proper Motion (Dec) Limit Flag
# COLUMN gaia_pm:        Gaia Total Proper Motion [mas/yr]
# COLUMN gaia_pmerr:     Gaia Total Proper Motion Unc. [mas/yr]
# COLUMN gaia_pmlim:     Gaia Total Proper Motion Limit Flag
# COLUMN st_radv:        Radial Velocity [km/s]
# COLUMN st_radverr1:    Radial Velocity Upper Unc. [km/s]
# COLUMN st_radverr2:    Radial Velocity Lower Unc. [km/s]
# COLUMN st_radvlim:     Radial Velocity Limit Flag
# COLUMN st_sp:          Spectral Type
# COLUMN st_spstr:       Spectral Type
# COLUMN st_sperr:       Spectral Type Unc.
# COLUMN st_splim:       Spectral Type Limit Flag
# COLUMN st_logg:        Stellar Surface Gravity [log10(cm/s**2)]
# COLUMN st_loggerr1:    Stellar Surface Gravity Upper Unc. [log10(cm/s**2)]
# COLUMN st_loggerr2:    Stellar Surface Gravity Lower Unc. [log10(cm/s**2)]
# COLUMN st_logglim:     Stellar Surface Gravity Limit Flag
# COLUMN st_lum:         Stellar Luminosity [log(Solar)]
# COLUMN st_lumerr1:     Stellar Luminosity Upper Unc. [log(Solar)]
# COLUMN st_lumerr2:     Stellar Luminosity Lower Unc. [log(Solar)]
# COLUMN st_lumlim:      Stellar Luminosity Limit Flag
# COLUMN st_dens:        Stellar Density [g/cm**3]
# COLUMN st_denserr1:    Stellar Density Upper Unc. [g/cm**3]
# COLUMN st_denserr2:    Stellar Density Lower Unc. [g/cm**3]
# COLUMN st_denslim:     Stellar Density Limit Flag
# COLUMN st_metfe:       Stellar Metallicity [dex]
# COLUMN st_metfeerr1:   Stellar Metallicity Upper Unc. [dex]
# COLUMN st_metfeerr2:   Stellar Metallicity Lower Unc. [dex]
# COLUMN st_metfelim:    Stellar Metallicity Limit Flag
# COLUMN st_metratio:    Metallicity Ratio
# COLUMN st_age:         Stellar Age [Gyr]
# COLUMN st_ageerr1:     Stellar Age Upper Unc. [Gyr]
# COLUMN st_ageerr2:     Stellar Age Lower Unc. [Gyr]
# COLUMN st_agelim:      Stellar Age Limit Flag
# COLUMN st_vsini:       Rot. Velocity V*sin(i) [km/s]
# COLUMN st_vsinierr1:   Rot. Velocity V*sin(i) Upper Unc. [km/s]
# COLUMN st_vsinierr2:   Rot. Velocity V*sin(i) Lower Unc. [km/s]
# COLUMN st_vsinilim:    Rot. Velocity V*sin(i) Limit Flag
# COLUMN st_acts:        Stellar Activity S-index
# COLUMN st_actserr:     Stellar Activity S-index Unc.
# COLUMN st_actslim:     Stellar Activity S-index Limit Flag
# COLUMN st_actr:        Stellar Activity log(R'HK)
# COLUMN st_actrerr:     Stellar Activity log(R'HK) Unc.
# COLUMN st_actrlim:     Stellar Activity log(R'HK) Limit Flag
# COLUMN st_actlx:       X-ray Activity log(L<sub>x</sub>)
# COLUMN st_actlxerr:    X-ray Activity log(L<sub>x</sub>) Unc.
# COLUMN st_actlxlim:    X-ray Activity log(L<sub>x</sub>) Limit Flag
# COLUMN swasp_id:       SWASP Identifier
# COLUMN st_nts:         Number of Time Series
# COLUMN st_nplc:        Number of Planet Transit Light Curves
# COLUMN st_nglc:        Number of General Light Curves
# COLUMN st_nrvc:        Number of Radial Velocity Time Series
# COLUMN st_naxa:        Number of Amateur Light Curves
# COLUMN st_nimg:        Number of Images
# COLUMN st_nspec:       Number of Spectra
# COLUMN st_uj:          U-band (Johnson) [mag]
# COLUMN st_ujerr:       U-band (Johnson) Unc. [mag]
# COLUMN st_ujlim:       U-band (Johnson) Limit Flag
# COLUMN st_vj:          V-band (Johnson) [mag]
# COLUMN st_vjerr:       V-band (Johnson) Unc. [mag]
# COLUMN st_vjlim:       V-band (Johnson) Limit Flag
# COLUMN st_bj:          B-band (Johnson) [mag]
# COLUMN st_bjerr:       B-band (Johnson) Unc. [mag]
# COLUMN st_bjlim:       B-band (Johnson) Limit Flag
# COLUMN st_rc:          R-band (Cousins) [mag]
# COLUMN st_rcerr:       R-band (Cousins) Unc. [mag]
# COLUMN st_rclim:       R-band (Cousins) Limit Flag
# COLUMN st_ic:          I-band (Cousins) [mag]
# COLUMN st_icerr:       I-band (Cousins) Unc. [mag]
# COLUMN st_iclim:       I-band (Cousins) Limit Flag
# COLUMN st_j:           J-band (2MASS) [mag]
# COLUMN st_jerr:        J-band (2MASS) Unc. [mag]
# COLUMN st_jlim:        J-band (2MASS) Limit Flag
# COLUMN st_h:           H-band (2MASS) [mag]
# COLUMN st_herr:        H-band (2MASS) Unc. [mag]
# COLUMN st_hlim:        H-band (2MASS) Limit Flag
# COLUMN st_k:           Ks-band (2MASS) [mag]
# COLUMN st_kerr:        Ks-band (2MASS) Unc. [mag]
# COLUMN st_klim:        Ks-band (2MASS) Limit Flag
# COLUMN st_wise1:       WISE 3.4um [mag]
# COLUMN st_wise1err:    WISE 3.4um Unc. [mag]
# COLUMN st_wise1lim:    WISE 3.4um Limit Flag
# COLUMN st_wise2:       WISE 4.6um [mag]
# COLUMN st_wise2err:    WISE 4.6um Unc. [mag]
# COLUMN st_wise2lim:    WISE 4.6um Limit Flag
# COLUMN st_wise3:       WISE 12.um [mag]
# COLUMN st_wise3err:    WISE 12.um Unc. [mag]
# COLUMN st_wise3lim:    WISE 12.um Limit Flag
# COLUMN st_wise4:       WISE 22.um [mag]
# COLUMN st_wise4err:    WISE 22.um Unc. [mag]
# COLUMN st_wise4lim:    WISE 22.um Limit Flag
# COLUMN st_irac1:       IRAC 3.6um [mag]
# COLUMN st_irac1err:    IRAC 3.6um Unc. [mag]
# COLUMN st_irac1lim:    IRAC 3.6um Limit Flag
# COLUMN st_irac2:       IRAC 4.5um [mag]
# COLUMN st_irac2err:    IRAC 4.5um Unc. [mag]
# COLUMN st_irac2lim:    IRAC 4.5um Limit Flag
# COLUMN st_irac3:       IRAC 5.8um [mag]
# COLUMN st_irac3err:    IRAC 5.8um Unc. [mag]
# COLUMN st_irac3lim:    IRAC 5.8um Limit Flag
# COLUMN st_irac4:       IRAC 8.0um [mag]
# COLUMN st_irac4err:    IRAC 8.0um Unc. [mag]
# COLUMN st_irac4lim:    IRAC 8.0um Limit Flag
# COLUMN st_mips1:       MIPS 24um [mag]
# COLUMN st_mips1err:    MIPS 24um Unc. [mag]
# COLUMN st_mips1lim:    MIPS 24um Limit Flag
# COLUMN st_mips2:       MIPS 70um [mag]
# COLUMN st_mips2err:    MIPS 70um Unc. [mag]
# COLUMN st_mips2lim:    MIPS 70um Limit Flag
# COLUMN st_mips3:       MIPS 160um [mag]
# COLUMN st_mips3err:    MIPS 160um Unc. [mag]
# COLUMN st_mips3lim:    MIPS 160um Limit Flag
# COLUMN st_iras1:       IRAS 12um Flux [Jy]
# COLUMN st_iras1err:    IRAS 12um Flux Unc. [Jy]
# COLUMN st_iras1lim:    IRAS 12um Flux Limit Flag
# COLUMN st_iras2:       IRAS 25um Flux [Jy]
# COLUMN st_iras2err:    IRAS 25um Flux Unc. [Jy]
# COLUMN st_iras2lim:    IRAS 25um Flux Limit Flag
# COLUMN st_iras3:       IRAS 60um Flux [Jy]
# COLUMN st_iras3err:    IRAS 60um Flux Unc. [Jy]
# COLUMN st_iras3lim:    IRAS 60um Flux Limit Flag
# COLUMN st_iras4:       IRAS 100um Flux [Jy]
# COLUMN st_iras4err:    IRAS 100um Flux Unc. [Jy]
# COLUMN st_iras4lim:    IRAS 100um Flux Limit Flag
# COLUMN st_photn:       Number of Photometry Measurements
# COLUMN st_umbj:        U-B (Johnson) [mag]
# COLUMN st_umbjerr:     U-B (Johnson) Unc. [mag]
# COLUMN st_umbjlim:     U-B (Johnson) Limit Flag
# COLUMN st_bmvj:        B-V (Johnson) [mag]
# COLUMN st_bmvjerr:     B-V (Johnson) Unc. [mag]
# COLUMN st_bmvjlim:     B-V (Johnson) Limit Flag
# COLUMN st_vjmic:       V-I (Johnson-Cousins) [mag]
# COLUMN st_vjmicerr:    V-I (Johnson-Cousins) Unc. [mag]
# COLUMN st_vjmiclim:    V-I (Johnson-Cousins) Limit Flag
# COLUMN st_vjmrc:       V-R (Johnson-Cousins) [mag]
# COLUMN st_vjmrcerr:    V-R (Johnson-Cousins) Unc. [mag]
# COLUMN st_vjmrclim:    V-R (Johnson-Cousins) Limit Flag
# COLUMN st_jmh2:        J-H (2MASS) [mag]
# COLUMN st_jmh2err:     J-H (2MASS) Unc. [mag]
# COLUMN st_jmh2lim:     J-H (2MASS) Limit Flag
# COLUMN st_hmk2:        H-Ks (2MASS) [mag]
# COLUMN st_hmk2err:     H-Ks (2MASS) Unc. [mag]
# COLUMN st_hmk2lim:     H-Ks (2MASS) Limit Flag
# COLUMN st_jmk2:        J-Ks (2MASS) [mag]
# COLUMN st_jmk2err:     J-Ks (2MASS) Unc. [mag]
# COLUMN st_jmk2lim:     J-Ks (2MASS) Limit Flag
# COLUMN st_bmy:         b-y (Stromgren) [mag]
# COLUMN st_bmyerr:      b-y (Stromgren) Unc. [mag]
# COLUMN st_bmylim:      b-y (Stromgren) Limit Flag
# COLUMN st_m1:          m1 (Stromgren) [mag]
# COLUMN st_m1err:       m1 (Stromgren) Unc. [mag]
# COLUMN st_m1lim:       m1 (Stromgren) Limit Flag
# COLUMN st_c1:          c1 (Stromgren) [mag]
# COLUMN st_c1err:       c1 (Stromgren) Unc. [mag]
# COLUMN st_c1lim:       c1 (Stromgren) Limit Flag
# COLUMN st_colorn:      Number of Color Measurements
*/

type ConfirmedExoPlanetCsvRowItem struct {
	// Host Name
	pl_hostname 	string

	// Planet Letter
	pl_letter		string

	// Planet Name
	pl_name string

	// Discovery Method
	pl_discmethod string

	// Number of Planets in System
	pl_pnum string

	// Orbital Period [days]
	pl_orbper string

	// Orbital Period Upper Unc. [days]
	pl_orbpererr1 string

	// Orbital Period Lower Unc. [days]
	pl_orbpererr2 string

	// Orbital Period Limit Flag
	pl_orbperlim string

	// Orbit Semi-Major Axis [AU])
	pl_orbsmax string

	// Orbit Semi-Major Axis Upper Unc. [AU]
	pl_orbsmaxerr1 string

	// Orbit Semi-Major Axis Lower Unc. [AU]
	pl_orbsmaxerr2 string

	// Orbit Semi-Major Axis Limit Flag
	pl_orbsmaxlim string

	// Eccentricity
	pl_orbeccen string

	// Eccentricity Upper Unc.
	pl_orbeccenerr1 string

	// Eccentricity Lower Unc.
	pl_orbeccenerr2 string

	// Eccentricity Limit Flag
	pl_orbeccenlim string

	// Inclination [deg]
	pl_orbincl string

	// Inclination Upper Unc. [deg]
	pl_orbinclerr1 string

	// Inclination Lower Unc. [deg]
	pl_orbinclerr2 string

	// Inclination Limit Flag
	pl_orbincllim string

	// Planet Mass or M*sin(i) [Jupiter mass]
	pl_bmassj string

	// Planet Mass or M*sin(i) Upper Unc. [Jupiter mass]
	pl_bmassjerr1 string

	// Planet Mass or M*sin(i) Lower Unc. [Jupiter mass]
	pl_bmassjerr2 string

	// Planet Mass or M*sin(i) Limit Flag
	pl_bmassjlim string

	// Planet Mass or M*sin(i) Provenance
	pl_bmassprov string

	// Planet Radius [Jupiter radii]
	pl_radj string

	// Planet Radius Upper Unc. [Jupiter radii]
	pl_radjerr1 string

	// Planet Radius Lower Unc. [Jupiter radii]
	pl_radjerr2 string

	// Planet Radius Limit Flag
	pl_radjlim string

	// Planet Density [g/cm**3]
	pl_dens string

	// Planet Density Upper Unc. [g/cm**3]
	pl_denserr1 string

	// Planet Density Lower Unc. [g/cm**3]
	pl_denserr2 string

	// Planet Density Limit Flag
	pl_denslim string

	// TTV Flag
	pl_ttvflag string

	// Kepler Field Flag
	pl_kepflag string

	// K2 Mission Flag
	pl_k2flag string

	// Number of Notes
	pl_nnotes string

	// RA [sexagesimal]
	ra_str string

	// RA [decimal degrees]
	ra string

	// Dec [sexagesimal]
	dec_str string

	// Dec [decimal degrees]
	dec string

	// Distance [pc]
	st_dist string

	// Distance Upper Unc. [pc]
	st_disterr1 string

	// Distance Lower Unc. [pc]
	st_disterr2 string

	// Distance Limit Flag
	st_distlim string

	// Optical Magnitude [mag]
	st_optmag string

	// Optical Magnitude Unc. [mag]
	st_optmagerr string

	// Optical Magnitude Limit Flag
	st_optmaglim string

	// Optical Magnitude Band
	st_optband string

	// G-band (Gaia) [mag]
	gaia_gmag string

	// G-band (Gaia) Unc. [mag]
	gaia_gmagerr string

	// G-band (Gaia) Limit Flag
	gaia_gmaglim string

	// Effective Temperature [K]
	st_teff string

	// Effective Temperature Upper Unc. [K]
	st_tefferr1 string

	// Effective Temperature Lower Unc. [K]
	st_tefferr2 string

	// Effective Temperature Limit Flag
	st_tefflim string

	// Stellar Mass [Solar mass]
	st_mass string

	// Stellar Mass Upper Unc. [Solar mass]
	st_masserr1 string

	// Stellar Mass Lower Unc. [Solar mass]
	st_masserr2 string

	// Stellar Mass Limit Flag
	st_masslim string

	// Stellar Radius [Solar radii]
	st_rad string

	// Stellar Radius Upper Unc. [Solar radii]
	st_raderr1 string

	// Stellar Radius Lower Unc. [Solar radii]
	st_raderr2 string

	// Stellar Radius Limit Flag
	st_radlim string

	// Date of Last Update
	rowupdate string

	// Planet Transit Flag
	pl_tranflag string

	// Planet RV Flag
	pl_rvflag string

	// Planet Imaging Flag
	pl_imgflag string

	// Planet Astrometry Flag
	pl_astflag string

	// Planet Orbital Modulation Flag
	pl_omflag string

	// Planet Circumbinary Flag
	pl_cbflag string

	// Calculated Angular Separation [mas]
	pl_angsep string

	// Calculated Angular Separation Upper Unc. [mas]
	pl_angseperr1 string

	// Calculated Angular Separation Lower Unc. [mas]
	pl_angseperr2 string

	// Time of Periastron [days]
	pl_orbtper string

	// Time of Periastron Upper Unc. [days]
	pl_orbtpererr1 string

	// Time of Periastron Lower Unc. [days]
	pl_orbtpererr2 string

	// Time of Periastron Limit Flag
	pl_orbtperlim string

	// Long. of Periastron [deg]
	pl_orblper string

	// Long. of Periastron Upper Unc. [deg]
	pl_orblpererr1 string

	// Long. of Periastron Lower Unc. [deg]
	pl_orblpererr2

	// Long. of Periastron Limit Flag
	pl_orblperlim

	// Radial Velocity Amplitude [m/s]
	pl_rvamp

	// Radial Velocity Amplitude Upper Unc. [m/s]
	pl_rvamperr1

	// Radial Velocity Amplitude Lower Unc. [m/s]
	pl_rvamperr2

	// Radial Velocity Amplitude Limit Flag
	pl_rvamplim

	// Equilibrium Temperature [K]
	pl_eqt

	// Equilibrium Temperature Upper Unc. [K]
	pl_eqterr1

	// Equilibrium Temperature Lower Unc. [K]
	pl_eqterr2

	// Equilibrium Temperature Limit Flag
	pl_eqtlim

	// Insolation Flux [Earth flux]
	pl_insol

	// Insolation Flux Upper Unc. [Earth flux]
	pl_insolerr1

	// Insolation Flux Lower Unc. [Earth flux]
	pl_insolerr2

	// Insolation Flux Limit Flag
	pl_insollim

	// Planet Mass [Jupiter mass]
	pl_massj

	// Planet Mass Upper Unc. [Jupiter mass]
	pl_massjerr1

	// Planet Mass Lower Unc. [Jupiter mass]
	pl_massjerr2

	// Planet Mass Limit Flag
	pl_massjlim

	// Planet M*sin(i) [Jupiter mass]
	pl_msinij

	// Planet M*sin(i) Upper Unc. [Jupiter mass]
	pl_msinijerr1

	// Planet M*sin(i) Lower Unc. [Jupiter mass]
	pl_msinijerr2

	// Planet M*sin(i) Limit Flag
	pl_msinijlim

	// Planet Mass [Earth mass]
	pl_masse

	// Planet Mass Upper Unc. [Earth mass]
	pl_masseerr1

:   Planet Mass Lower Unc. [Earth mass]
pl_masseerr2

:    Planet Mass Limit Flag
pl_masselim

:      Planet M*sin(i) [Earth mass]
pl_msinie

:  Planet M*sin(i) Upper Unc. [Earth mass]
pl_msinieerr1

:  Planet M*sin(i) Lower Unc. [Earth mass]
pl_msinieerr2


pl_msinielim:   Planet M*sin(i) Limit Flag


pl_bmasse:      Planet Mass or M*sin(i) [Earth mass]


pl_bmasseerr1:  Planet Mass or M*sin(i) Upper Unc. [Earth mass]


pl_bmasseerr2:  Planet Mass or M*sin(i) Lower Unc. [Earth mass]


pl_bmasselim:   Planet Mass or M*sin(i) Limit Flag


pl_rade:        Planet Radius [Earth radii]


pl_radeerr1:    Planet Radius Upper Unc. [Earth radii]


pl_radeerr2:    Planet Radius Lower Unc. [Earth radii]


pl_radelim:     Planet Radius Limit Flag


pl_rads:        Planet Radius [Solar radii]


pl_radserr1:    Planet Radius Upper Unc. [Solar radii]


pl_radserr2:    Planet Radius Lower Unc. [Solar radii]


pl_radslim:     Planet Radius Limit Flag


pl_trandep:     Transit Depth [percent]


pl_trandeperr1: Transit Depth Upper Unc. [percent]


pl_trandeperr2: Transit Depth Lower Unc. [percent]


pl_trandeplim:  Transit Depth Limit Flag


pl_trandur:     Transit Duration [days]


pl_trandurerr1: Transit Duration Upper Unc. [days]


pl_trandurerr2: Transit Duration Lower Unc. [days]


pl_trandurlim:  Transit Duration Limit Flag


pl_tranmid:     Transit Midpoint [days]


pl_tranmiderr1: Transit Midpoint Upper Unc. [days]


pl_tranmiderr2: Transit Midpoint Lower Unc. [days]


pl_tranmidlim:  Transit Midpoint Limit Flag


pl_tsystemref:  Time System Reference


pl_imppar:      Impact Parameter


pl_impparerr1:  Impact Parameter Upper Unc.


pl_impparerr2:  Impact Parameter Lower Unc.


pl_impparlim:   Impact Parameter Limit Flag


pl_occdep:      Occultation Depth [percentage]


pl_occdeperr1:  Occultation Depth Upper Unc. [percentage]


pl_occdeperr2:  Occultation Depth Lower Unc. [percentage]


pl_occdeplim:   Occultation Depth Limit Flag


pl_ratdor:      Ratio of Distance to Stellar Radius


pl_ratdorerr1:  Ratio of Distance to Stellar Radius Upper Unc.


pl_ratdorerr2:  Ratio of Distance to Stellar Radius Lower Unc.


pl_ratdorlim:   Ratio of Distance to Stellar Radius Limit Flag


pl_ratror:      Ratio of Planet to Stellar Radius


pl_ratrorerr1:  Ratio of Planet to Stellar Radius Upper Unc.


pl_ratrorerr2:  Ratio of Planet to Stellar Radius Lower Unc.


pl_ratrorlim:   Ratio of Planet to Stellar Radius Limit Flag


pl_def_reflink: Default Reference


pl_disc:        Year of Discovery


pl_disc_reflink: Discovery Reference


pl_locale:      Discovery Locale


pl_facility:    Discovery Facility


pl_telescope:   Discovery Telescope


pl_instrument:  Discovery Instrument


pl_status:      Status


pl_mnum:        Number of Moons in System


pl_st_npar:     Number of Stellar and Planet Parameters


pl_st_nref:     Number of Stellar and Planet References


pl_pelink:      Link to Exoplanet Encyclopaedia


pl_edelink:     Link to Exoplanet Data Explorer


pl_publ_date:   Publication Date


hd_name:        HD Name


hip_name:       HIP Name


st_rah:         RA [hrs]


st_glon:        Galactic Longitude [deg]


st_glat:        Galactic Latitude [deg]


st_elon:        Ecliptic Longitude [deg]


st_elat:        Ecliptic Latitude [deg]



st_plx:         Parallax [mas]


st_plxerr1:     Parallax Upper Unc. [mas]


st_plxerr2:     Parallax Lower Unc. [mas]


st_plxlim:      Parallax Limit Flag


gaia_plx:       Gaia Parallax [mas]


gaia_plxerr1:   Gaia Parallax Upper Unc. [mas]


gaia_plxerr2:   Gaia Parallax Lower Unc. [mas]


gaia_plxlim:    Gaia Parallax Limit Flag


gaia_dist:      Gaia Distance [pc]


gaia_disterr1:  Gaia Distance Upper Unc. [pc]


gaia_disterr2:  Gaia Distance Lower Unc. [pc]


gaia_distlim:   Gaia Distance Limit Flag


st_pmra:        Proper Motion (RA) [mas/yr]


st_pmraerr:     Proper Motion (RA) Unc. [mas/yr]


st_pmralim:     Proper Motion (RA) Limit Flag


st_pmdec:       Proper Motion (Dec) [mas/yr]


st_pmdecerr:    Proper Motion (Dec) Unc. [mas/yr]


st_pmdeclim:    Proper Motion (Dec) Limit Flag


st_pm:          Total Proper Motion [mas/yr]


st_pmerr:       Total Proper Motion Unc. [mas/yr]


st_pmlim:       Total Proper Motion Limit Flag


gaia_pmra:      Gaia Proper Motion (RA) [mas/yr]


gaia_pmraerr:   Gaia Proper Motion (RA) Unc. [mas/yr]


gaia_pmralim:   Gaia Proper Motion (RA) Limit Flag


gaia_pmdec:     Gaia Proper Motion (Dec) [mas/yr]


gaia_pmdecerr:  Gaia Proper Motion (Dec) Unc. [mas/yr]


gaia_pmdeclim:  Gaia Proper Motion (Dec) Limit Flag


gaia_pm:        Gaia Total Proper Motion [mas/yr]


gaia_pmerr:     Gaia Total Proper Motion Unc. [mas/yr]


gaia_pmlim:     Gaia Total Proper Motion Limit Flag


st_radv:        Radial Velocity [km/s]


st_radverr1:    Radial Velocity Upper Unc. [km/s]


st_radverr2:    Radial Velocity Lower Unc. [km/s]


st_radvlim:     Radial Velocity Limit Flag


st_sp:          Spectral Type


st_spstr:       Spectral Type


st_sperr:       Spectral Type Unc.


st_splim:       Spectral Type Limit Flag


st_logg:        Stellar Surface Gravity [log10(cm/s**2)]


st_loggerr1:    Stellar Surface Gravity Upper Unc. [log10(cm/s**2)]


st_loggerr2:    Stellar Surface Gravity Lower Unc. [log10(cm/s**2)]


st_logglim:     Stellar Surface Gravity Limit Flag


st_lum:         Stellar Luminosity [log(Solar)]


st_lumerr1:     Stellar Luminosity Upper Unc. [log(Solar)]


st_lumerr2:     Stellar Luminosity Lower Unc. [log(Solar)]


st_lumlim:      Stellar Luminosity Limit Flag


st_dens:        Stellar Density [g/cm**3]


st_denserr1:    Stellar Density Upper Unc. [g/cm**3]


st_denserr2:    Stellar Density Lower Unc. [g/cm**3]


st_denslim:     Stellar Density Limit Flag


st_metfe:       Stellar Metallicity [dex]


st_metfeerr1:   Stellar Metallicity Upper Unc. [dex]


st_metfeerr2:   Stellar Metallicity Lower Unc. [dex]


st_metfelim:    Stellar Metallicity Limit Flag


st_metratio:    Metallicity Ratio


st_age:         Stellar Age [Gyr]


st_ageerr1:     Stellar Age Upper Unc. [Gyr]


st_ageerr2:     Stellar Age Lower Unc. [Gyr]


st_agelim:      Stellar Age Limit Flag


st_vsini:       Rot. Velocity V*sin(i) [km/s]


st_vsinierr1:   Rot. Velocity V*sin(i) Upper Unc. [km/s]


st_vsinierr2:   Rot. Velocity V*sin(i) Lower Unc. [km/s]


st_vsinilim:    Rot. Velocity V*sin(i) Limit Flag


st_acts:        Stellar Activity S-index


st_actserr:     Stellar Activity S-index Unc.


st_actslim:     Stellar Activity S-index Limit Flag


st_actr:        Stellar Activity log(R'HK)


st_actrerr:     Stellar Activity log(R'HK) Unc.


st_actrlim:     Stellar Activity log(R'HK) Limit Flag


st_actlx:       X-ray Activity log(L<sub>x</sub>)


st_actlxerr:    X-ray Activity log(L<sub>x</sub>) Unc.


st_actlxlim:    X-ray Activity log(L<sub>x</sub>) Limit Flag


swasp_id:       SWASP Identifier


st_nts:         Number of Time Series


st_nplc:        Number of Planet Transit Light Curves


st_nglc:        Number of General Light Curves


st_nrvc:        Number of Radial Velocity Time Series


st_naxa:        Number of Amateur Light Curves


st_nimg:        Number of Images


st_nspec:       Number of Spectra


st_uj:          U-band (Johnson) [mag]


st_ujerr:       U-band (Johnson) Unc. [mag]


st_ujlim:       U-band (Johnson) Limit Flag


st_vj:          V-band (Johnson) [mag]


st_vjerr:       V-band (Johnson) Unc. [mag]


st_vjlim:       V-band (Johnson) Limit Flag


st_bj:          B-band (Johnson) [mag]


st_bjerr:       B-band (Johnson) Unc. [mag]


st_bjlim:       B-band (Johnson) Limit Flag


st_rc:          R-band (Cousins) [mag]


st_rcerr:       R-band (Cousins) Unc. [mag]


st_rclim:       R-band (Cousins) Limit Flag


st_ic:          I-band (Cousins) [mag]


st_icerr:       I-band (Cousins) Unc. [mag]


st_iclim:       I-band (Cousins) Limit Flag


st_j:           J-band (2MASS) [mag]


st_jerr:        J-band (2MASS) Unc. [mag]


st_jlim:        J-band (2MASS) Limit Flag


st_h:           H-band (2MASS) [mag]


st_herr:        H-band (2MASS) Unc. [mag]


st_hlim:        H-band (2MASS) Limit Flag


st_k:           Ks-band (2MASS) [mag]


st_kerr:        Ks-band (2MASS) Unc. [mag]


st_klim:        Ks-band (2MASS) Limit Flag


st_wise1:       WISE 3.4um [mag]


st_wise1err:    WISE 3.4um Unc. [mag]


st_wise1lim:    WISE 3.4um Limit Flag


st_wise2:       WISE 4.6um [mag]


st_wise2err:    WISE 4.6um Unc. [mag]


st_wise2lim:    WISE 4.6um Limit Flag


st_wise3:       WISE 12.um [mag]


st_wise3err:    WISE 12.um Unc. [mag]


st_wise3lim:    WISE 12.um Limit Flag


st_wise4:       WISE 22.um [mag]


st_wise4err:    WISE 22.um Unc. [mag]


st_wise4lim:    WISE 22.um Limit Flag


st_irac1:       IRAC 3.6um [mag]


st_irac1err:    IRAC 3.6um Unc. [mag]


st_irac1lim:    IRAC 3.6um Limit Flag


st_irac2:       IRAC 4.5um [mag]


st_irac2err:    IRAC 4.5um Unc. [mag]


st_irac2lim:    IRAC 4.5um Limit Flag


st_irac3:       IRAC 5.8um [mag]


st_irac3err:    IRAC 5.8um Unc. [mag]


st_irac3lim:    IRAC 5.8um Limit Flag


st_irac4:       IRAC 8.0um [mag]


st_irac4err:    IRAC 8.0um Unc. [mag]


st_irac4lim:    IRAC 8.0um Limit Flag


st_mips1:       MIPS 24um [mag]


st_mips1err:    MIPS 24um Unc. [mag]


st_mips1lim:    MIPS 24um Limit Flag


st_mips2:       MIPS 70um [mag]


st_mips2err:    MIPS 70um Unc. [mag]


st_mips2lim:    MIPS 70um Limit Flag


st_mips3:       MIPS 160um [mag]


st_mips3err:    MIPS 160um Unc. [mag]


st_mips3lim:    MIPS 160um Limit Flag


st_iras1:       IRAS 12um Flux [Jy]


st_iras1err:    IRAS 12um Flux Unc. [Jy]


st_iras1lim:    IRAS 12um Flux Limit Flag


st_iras2:       IRAS 25um Flux [Jy]


st_iras2err:    IRAS 25um Flux Unc. [Jy]


st_iras2lim:    IRAS 25um Flux Limit Flag


st_iras3:       IRAS 60um Flux [Jy]


st_iras3err:    IRAS 60um Flux Unc. [Jy]


st_iras3lim:    IRAS 60um Flux Limit Flag


st_iras4:       IRAS 100um Flux [Jy]


st_iras4err:    IRAS 100um Flux Unc. [Jy]


st_iras4lim:    IRAS 100um Flux Limit Flag


st_photn:       Number of Photometry Measurements


st_umbj:        U-B (Johnson) [mag]


st_umbjerr:     U-B (Johnson) Unc. [mag]


st_umbjlim:     U-B (Johnson) Limit Flag


st_bmvj:        B-V (Johnson) [mag]


st_bmvjerr:     B-V (Johnson) Unc. [mag]


st_bmvjlim:     B-V (Johnson) Limit Flag


st_vjmic:       V-I (Johnson-Cousins) [mag]


st_vjmicerr:    V-I (Johnson-Cousins) Unc. [mag]


st_vjmiclim:    V-I (Johnson-Cousins) Limit Flag


st_vjmrc:       V-R (Johnson-Cousins) [mag]


st_vjmrcerr:    V-R (Johnson-Cousins) Unc. [mag]


st_vjmrclim:    V-R (Johnson-Cousins) Limit Flag


st_jmh2:        J-H (2MASS) [mag]


st_jmh2err:     J-H (2MASS) Unc. [mag]


st_jmh2lim:     J-H (2MASS) Limit Flag


st_hmk2:        H-Ks (2MASS) [mag]


st_hmk2err:     H-Ks (2MASS) Unc. [mag]


st_hmk2lim:     H-Ks (2MASS) Limit Flag


st_jmk2:        J-Ks (2MASS) [mag]

:     J-Ks (2MASS) Unc. [mag]
st_jmk2err

:     J-Ks (2MASS) Limit Flag
st_jmk2lim

:         b-y (Stromgren) [mag]
st_bmy

:      b-y (Stromgren) Unc. [mag]
st_bmyerr

:      b-y (Stromgren) Limit Flag
st_bmylim

:          m1 (Stromgren) [mag]
st_m1

:       m1 (Stromgren) Unc. [mag]
st_m1err

:       m1 (Stromgren) Limit Flag
st_m1lim

:          c1 (Stromgren) [mag]
st_c1

:       c1 (Stromgren) Unc. [mag]
st_c1err

:       c1 (Stromgren) Limit Flag
st_c1lim

:      Number of Color Measurements
st_colorn


}